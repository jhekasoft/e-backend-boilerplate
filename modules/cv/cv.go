package cv

import (
	"e-backend/modules/cv/handler"
	"e-backend/modules/cv/repository"
	"e-backend/modules/cv/service"
	internalModels "e-backend/pkg/ebackend/models"
	"path"
)

const CVBaseURL = "/"
const CVDataPath = "./modules/cv/data"
const CVTemplatesPath = "./modules/cv/templates"

type CVModule struct {
}

func (m *CVModule) Name() string {
	return "CV"
}

func (m *CVModule) Run(c *internalModels.Core) error {
	repo := repository.NewRepository(CVDataPath)
	services := service.NewService(repo, CVBaseURL, CVTemplatesPath)
	h := handler.NewHandler(services)

	c.Echo.GET("/cv/developer-timeline", h.GetDevTimeline)
	c.Echo.GET("/cv/cv", h.GetCV)
	c.Echo.Static("/cv/public", path.Join(CVDataPath, "public"))
	c.Echo.GET("/cv/latex", h.GetCVLatex)

	return nil
}

func NewModule() internalModels.Module {
	return &CVModule{}
}
