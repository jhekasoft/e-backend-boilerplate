package handler

import (
	"e-backend-boilerplate/modules/smarthome/models"
	"e-backend-boilerplate/modules/smarthome/service"
	"errors"
	"fmt"
	"regexp"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/labstack/gommon/log"
)

const SensorMeteo = "meteo"

// type SensorValueRequest struct {
// 	Temperature *float32 `json:"temperature"`
// 	Humidity    *float32 `json:"humidity"`
// 	Pressure    *float32 `json:"pressure"`
// }

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) MQTTSensorValue(client mqtt.Client, msg mqtt.Message) {
	name1, name2, name3, err := h.parseSensorValueTopic(msg.Topic())
	if err != nil {
		log.Error(err)
		return
	}

	item := models.SmartHomeSensorValue{
		Name1:  name1,
		Name2:  name2,
		Name3:  name3,
		Sensor: SensorMeteo,
	}
	item.Value.Set(string(msg.Payload()))

	_, err = h.service.Create(item)
	if err != nil {
		log.Error(err)
		return
	}

	// Logging
	log.Info(fmt.Sprintf("MQTT in [%s] %s\n", msg.Topic(), string(msg.Payload())))
}

func (h *Handler) parseSensorValueTopic(topic string) (name1, name2, name3 string, err error) {
	r := regexp.MustCompile(`[\w\d]+\/([\w\d]+)\/([\w\d]+)\/([\w\d]+)\/`)
	matches := r.FindStringSubmatch(topic)
	if len(matches) < 4 {
		err = errors.New("no names in MQTT topics")
		return
	}
	name1 = matches[1]
	name2 = matches[2]
	name3 = matches[3]

	return
}
