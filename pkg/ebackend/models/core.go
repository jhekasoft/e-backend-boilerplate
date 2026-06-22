package models

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Core struct {
	Version   string
	BuildTime string
	Config    Config
	Echo      *echo.Echo
	MQTT      *mqtt.Client
	DB        *gorm.DB
	Trans     *ut.Translator
}
