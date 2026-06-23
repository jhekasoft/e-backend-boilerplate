package auth

import (
	"e-backend-boilerplate/modules/auth/handler"
	"e-backend-boilerplate/modules/auth/models"
	"e-backend-boilerplate/modules/auth/repository"
	"e-backend-boilerplate/modules/auth/service"

	"github.com/jhekasoft/e-backend/http/middleware"
	internalModels "github.com/jhekasoft/e-backend/models"
)

type AuthModule struct {
}

func (m *AuthModule) Name() string {
	return "Auth"
}

func (m *AuthModule) Run(c *internalModels.Core) error {
	c.DB.AutoMigrate(&models.User{})

	repo := repository.NewRepository(c.DB)
	services := service.NewService(repo, c.Config.Auth.JWTSecretKey)
	h := handler.NewHandler(services)

	authMiddleware := middleware.AuthMiddleware(c.Config.Auth.JWTSecretKey)

	c.Echo.POST("/auth/users", h.CreateItem) // Registration
	c.Echo.POST("/auth/signin", h.SignIn)
	c.Echo.GET("/auth/users/me", h.CurrentUser, authMiddleware)

	return nil
}

func NewModule() internalModels.Module {
	return &AuthModule{}
}
