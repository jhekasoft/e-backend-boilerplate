package admins

import (
	"e-backend/modules/admins/handler"
	"e-backend/modules/admins/models"
	"e-backend/modules/admins/repository"
	"e-backend/modules/admins/service"
	internalModels "e-backend/pkg/ebackend/models"
)

type AdminsModule struct {
}

func (m *AdminsModule) Name() string {
	return "Admins"
}

func (m *AdminsModule) Run(c *internalModels.Core) error {
	c.DB.AutoMigrate(&models.Admin{})

	repo := repository.NewRepository(c.DB)
	services := service.NewService(repo)
	h := handler.NewHandler(services)

	c.Echo.GET("/admins", h.GetList)
	c.Echo.GET("/admins/:id", h.GetItem)
	c.Echo.POST("/admins", h.CreateItem)
	c.Echo.PUT("/admins/:id", h.UpdateItem)
	c.Echo.DELETE("/admins/:id", h.DeleteItem)

	return nil
}

func NewModule() internalModels.Module {
	return &AdminsModule{}
}
