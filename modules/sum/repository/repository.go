package repository

import (
	"e-backend-boilerplate/modules/sum/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Paginate(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

func (r *Repository) GetMany(filter models.ArticleListFilter) (items []models.Article, err error) {
	err = r.getListQuery(filter).
		Scopes(r.Paginate(filter.GetOffset(), filter.GetLimit())).
		Find(&items).
		Error
	return
}

func (r *Repository) GetTotal(filter models.ArticleListFilter) (count int64, err error) {
	err = r.getListQuery(filter).Count(&count).Error
	return
}

func (r *Repository) getListQuery(filter models.ArticleListFilter) *gorm.DB {
	var item models.Article
	tx := r.db.Model(&item)
	tx.Where("type = ?", "article")
	tx.Order("rank DESC")

	if len(filter.Search) > 0 {
		tx.Select("*, ts_rank_cd(to_tsvector('ukrainian', title), to_tsquery('ukrainian', ?)) AS rank", filter.Search)
		tx.Where("to_tsvector('ukrainian', title) @@ to_tsquery('ukrainian', ?)", filter.Search)
	}

	return tx
}

func (r *Repository) CreateInBatches(items []models.Article) (affected int64, err error) {
	result := r.db.CreateInBatches(items, 1000)
	err = result.Error
	affected = result.RowsAffected
	return
}
