package handler

import (
	"e-backend-boilerplate/modules/tamagluchi/models"
	"e-backend-boilerplate/modules/tamagluchi/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		service *service.Service
	}

	CreatePetRequest struct {
		Name string `validate:"required"`
		Type string `validate:"required,oneof=cat dog"`
	}

	CreatePetResponse struct {
		State models.TamagluchiState
	}

	CalculateRequest struct {
		State  models.TamagluchiState `validate:"required"`
		Period int                    `validate:"required,min=1"` // in seconds
	}

	CalculateResponse struct {
		State models.TamagluchiState
	}
)

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Create(c echo.Context) error {
	req := new(CreatePetRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	pet := models.Pet{
		Name: req.Name,
		Type: req.Type,
	}

	state, err := h.service.CreatePet(pet)
	if err != nil {
		return err
	}

	resp := CreatePetResponse{
		State: state,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) Calculate(c echo.Context) error {
	req := new(CalculateRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	state, err := h.service.Calculate(req.State, req.Period)
	if err != nil {
		return err
	}

	resp := CalculateResponse{
		State: state,
	}

	return c.JSON(http.StatusOK, resp)
}
