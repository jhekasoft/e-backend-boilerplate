package smarthome

import (
	"e-backend-boilerplate/modules/smarthome/handler"
	"e-backend-boilerplate/modules/smarthome/models"
	"e-backend-boilerplate/modules/smarthome/repository"
	"e-backend-boilerplate/modules/smarthome/service"
	internalModels "e-backend-boilerplate/pkg/ebackend/models"
)

type SmartHomeModule struct {
}

func (m *SmartHomeModule) Name() string {
	return "SmartHome"
}

func (m *SmartHomeModule) Run(c *internalModels.Core) error {
	c.DB.AutoMigrate(&models.SmartHomeSensorValue{})

	repo := repository.NewRepository(c.DB)
	services := service.NewService(repo)
	h := handler.NewHandler(services)

	if c.MQTT != nil {
		(*c.MQTT).Subscribe("smarthome/+/+/+/sensor/meteo", 0, h.MQTTSensorValue)
	}

	return nil
}

func NewModule() internalModels.Module {
	return &SmartHomeModule{}
}
