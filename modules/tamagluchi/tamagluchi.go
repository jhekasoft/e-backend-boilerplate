package tamagluchi

import (
	"e-backend/modules/tamagluchi/handler"
	"e-backend/modules/tamagluchi/repository"
	"e-backend/modules/tamagluchi/service"
	internalModels "e-backend/pkg/ebackend/models"
)

type TamagluchiModule struct {
}

func (m *TamagluchiModule) Name() string {
	return "Tamagluchi"
}

func (m *TamagluchiModule) Run(c *internalModels.Core) error {
	repo := repository.NewRepository(c.DB)
	services := service.NewService(repo)
	h := handler.NewHandler(services)

	c.Echo.POST("/tamagluchi", h.Create)
	c.Echo.POST("/tamagluchi/calculate", h.Calculate)

	return nil
}

func NewModule() internalModels.Module {
	return &TamagluchiModule{}
}
