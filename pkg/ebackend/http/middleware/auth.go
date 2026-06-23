package middleware

import (
	"e-backend-boilerplate/pkg/ebackend/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(jwtSecret string) echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.AuthClaims)
		},
		SigningKey: []byte(jwtSecret),
	}

	return echojwt.WithConfig(config)
}
