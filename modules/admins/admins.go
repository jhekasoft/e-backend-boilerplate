package admins

import (
	"e-backend-boilerplate/modules/admins/handler"
	"e-backend-boilerplate/modules/admins/models"
	"e-backend-boilerplate/modules/admins/repository"
	"e-backend-boilerplate/modules/admins/service"

	internalModels "github.com/jhekasoft/e-backend/models"
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
