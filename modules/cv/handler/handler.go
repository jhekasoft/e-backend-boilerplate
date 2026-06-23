package handler

import (
	"e-backend-boilerplate/modules/cv/models"
	"e-backend-boilerplate/modules/cv/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		service *service.Service
	}

	DevTimelineFilter struct {
		Reverse bool `query:"reverse"`
	}

	DevTimelineResponse struct {
		Data []models.DevTimelineItem `json:"data"`
	}

	CVResponse struct {
		Data CVData `json:"data"`
	}

	CVData struct {
		Common           models.CVCommon            `json:"common"`
		Education        []models.CVEducationItem   `json:"education"`
		Experience       []models.CVExperienceItem  `json:"experience"`
		Publications     []models.CVPublication     `json:"publications"`
		SoftwareProjects []models.CVSoftwareProject `json:"softwareProjects"`
	}
)

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetDevTimeline(c echo.Context) error {
	req := new(DevTimelineFilter)
	err := c.Bind(req)
	if err != nil {
		return err
	}

	list, err := h.service.GetDevTimeline(req.Reverse)
	if err != nil {
		return err
	}

	resp := DevTimelineResponse{
		Data: list,
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCV(c echo.Context) error {
	common, err := h.service.GetCVCommon()
	if err != nil {
		return err
	}

	education, err := h.service.GetCVEducation()
	if err != nil {
		return err
	}

	experience, err := h.service.GetCVExperience()
	if err != nil {
		return err
	}

	publications, err := h.service.GetCVPublications()
	if err != nil {
		return err
	}

	softwareProjects, err := h.service.GetCVSoftwareProjects()
	if err != nil {
		return err
	}

	resp := CVResponse{
		Data: CVData{
			Common:           *common,
			Education:        education,
			Experience:       experience,
			Publications:     publications,
			SoftwareProjects: softwareProjects,
		},
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCVLatex(c echo.Context) error {
	latexContent, err := h.service.GetCVLatex()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, latexContent)
}
