package service

import (
	"e-backend/modules/sum/models"
	"e-backend/modules/sum/repository"
	"e-backend/pkg/ebackend/crud"
	"log"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetManyWithTotal(filter models.ArticleListFilter) (items []models.Article, total int64, err error) {
	// Filter acute accents
	filter.Search = s.filterWord(filter.Search)

	items, err = s.repo.GetMany(filter)
	if err != nil {
		return
	}

	total, err = s.repo.GetTotal(filter)
	return
}

func (s *Service) GetWordOrAlternatives(title string) (item *models.Article, alts []string, err error) {
	// Always slice (not nil)
	alts = make([]string, 0)

	// Filter acute accents
	title = s.filterWord(title)

	items, err := s.repo.GetMany(models.ArticleListFilter{
		ListFilter: crud.ListFilter{Limit: 10, Offset: 0},
		Search:     title,
	})
	if err != nil {
		return
	}

	if len(items) < 1 {
		return
	}

	// Search exact word
	for _, it := range items {
		if it.Title != nil && strings.EqualFold(*it.Title, title) {
			item = &it
			return
		}
	}

	// Alternatives
	for _, it := range items {
		if it.Word == nil {
			continue
		}
		alts = append(alts, *it.Word)
	}

	return
}

func (s *Service) filterWord(word string) string {
	// Replace acute accents
	// Example: АБОНУВА́ТИСЯ -> АБОНУВАТИСЯ
	return strings.Replace(word, "\u0301", "", -1)
}

func (s *Service) Import() (affected int64, err error) {
	dbPath := "./modules/sum/data/import.db"

	var importItems []models.ImportLink
	importDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return
	}

	q := importDB.Where("desc IS NOT NULL AND type = ?", "article")
	err = q.Find(&importItems).Error
	if err != nil {
		return
	}

	var insertItems []models.Article
	for i, item := range importItems {
		log.Printf("Link %d: %s\n", i, item.URL)

		insertItems = append(insertItems, models.Article{
			Model: gorm.Model{ID: item.ID},
			Type:  item.Type,
			Word:  item.Word,
			Desc:  item.Desc,
			Title: item.Title,
		})
	}

	affected, err = s.repo.CreateInBatches(insertItems)
	return
}
