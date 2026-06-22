package ebackend

import (
	internalHttp "e-backend/pkg/ebackend/http"
	"e-backend/pkg/ebackend/models"
	"fmt"
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/go-playground/locales/uk"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	uk_translations "github.com/go-playground/validator/v10/translations/uk"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

type HTTPErrorResponse struct {
	Message  string
	Messages map[string]string `json:",omitempty"`
}

type HTTPApp struct {
	Core models.Core
}

func NewHTTPApp(config models.Config) HTTPApp {
	return HTTPApp{Core: models.Core{
		Config: config,
	}}
}

func (a *HTTPApp) Run(modules []models.Module, version, buildTime string) {
	a.Core.Version = version
	a.Core.BuildTime = buildTime

	config := a.Core.Config

	// Connect to the database
	dbLogLevel := logger.Error
	if config.IsDevelop() {
		dbLogLevel = logger.Info
	}
	db, err := gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(dbLogLevel),
	})
	if err != nil {
		log.Fatalf("Database connection error: %v\n", err)
	}
	a.Core.DB = db

	// Prepare translator
	uk := uk.New()
	uni := ut.New(uk, uk)
	trans, _ := uni.GetTranslator("uk")
	validate := validator.New(validator.WithRequiredStructEnabled())
	uk_translations.RegisterDefaultTranslations(validate, trans)
	a.Core.Trans = &trans

	// Prepare MQTT
	if config.MQTT.Enabled {
		var mqttConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
			fmt.Println("MQTT connected")
		}
		var mqttConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
			fmt.Printf("MQTT connection lost: %v", err)
		}

		mqttOpts := mqtt.NewClientOptions()
		mqttOpts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.MQTT.Host, config.MQTT.Port))
		mqttOpts.SetClientID("e-backend_client")
		mqttOpts.OnConnect = mqttConnectHandler
		mqttOpts.OnConnectionLost = mqttConnectLostHandler
		mqttClient := mqtt.NewClient(mqttOpts)
		if mqttToken := mqttClient.Connect(); mqttToken.Wait() && mqttToken.Error() != nil {
			log.Fatalf("MQTT init error: %v\n", mqttToken.Error())
		}
		a.Core.MQTT = &mqttClient
	}

	// Prepare HTTP-server
	a.Core.Echo = echo.New()
	a.Core.Echo.HideBanner = true
	a.Core.Echo.HTTPErrorHandler = a.httpErrorHandler

	a.Core.Echo.Validator = &CustomValidator{
		validator: validate,
	}

	// a.Core.Echo.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	a.Core.Echo.Use(middleware.Logger())
	a.Core.Echo.Use(middleware.Recover())

	a.Core.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		ExposeHeaders:    []string{"Server", "Content-Type", "Content-Disposition"},
		AllowCredentials: true,
	}))

	// Run modules
	for _, m := range modules {
		fmt.Printf("Run module %s\n", m.Name())
		err := m.Run(&a.Core)
		if err != nil {
			log.Fatalf("Module run error: %v\n", err)
		}
	}

	// Run HTTP-server
	a.Core.Echo.Logger.Fatal(a.Core.Echo.Start(fmt.Sprintf(":%d", config.HTTP.Port)))
}

func (a *HTTPApp) httpErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	var (
		code = http.StatusInternalServerError
		msg  HTTPErrorResponse
	)

	switch e := err.(type) {
	case *echo.HTTPError:
		code = e.Code
		msg = HTTPErrorResponse{Message: fmt.Sprintf("%v", e.Message)}
	case validator.ValidationErrors:
		code = http.StatusBadRequest
		messages := make(map[string]string)
		for _, err := range e {
			var message string
			if a.Core.Trans != nil {
				message = err.Translate(*a.Core.Trans)
			} else {
				message = err.Error()
			}
			messages[err.Field()] = message
		}
		msg = HTTPErrorResponse{Message: "Помилка валідації форми", Messages: messages}
	case *internalHttp.CustomValidationError:
		code = http.StatusBadRequest
		msg = HTTPErrorResponse{Message: e.Error(), Messages: e.Messages}
	default:
		msg = HTTPErrorResponse{Message: err.Error()}
	}

	// Send response
	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(code)
	} else {
		err = c.JSON(code, msg)
	}
	if err != nil {
		a.Core.Echo.Logger.Error(err)
	}
}
