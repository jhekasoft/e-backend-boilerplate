package sum

import (
	"e-backend/modules/sum/handler"
	"e-backend/modules/sum/models"
	"e-backend/modules/sum/repository"
	"e-backend/modules/sum/service"
	internalModels "e-backend/pkg/ebackend/models"
	"log"
)

type SumModule struct {
}

func (m *SumModule) Name() string {
	return "Sum"
}

func (m *SumModule) Run(c *internalModels.Core) error {
	// Migrations
	c.DB.AutoMigrate(&models.Article{})
	result := c.DB.Exec(`CREATE INDEX IF NOT EXISTS "search_idx"
ON "sum_articles" USING GIN (to_tsvector('ukrainian', title))`)
	if result.Error != nil {
		log.Println(result.Error)
	}

	repo := repository.NewRepository(c.DB)
	services := service.NewService(repo)
	h := handler.NewHandler(services)

	c.Echo.GET("/sum/articles", h.GetList)
	c.Echo.GET("/sum/articles/:word", h.GetWord)
	c.Echo.POST("/sum/import", h.Import) // TODO: move to the cmd

	return nil
}

func NewModule() internalModels.Module {
	return &SumModule{}
}
