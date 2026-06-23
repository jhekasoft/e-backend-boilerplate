package models

import (
	"e-backend-boilerplate/pkg/ebackend/crud"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	// URL        string `gorm:"uniqueIndex"`
	Type string  `gorm:"index"`
	Word *string `gorm:"index"`
	// ParentID   *uint
	// ParentURL  *string
	// Vocabulary string `gorm:"default:sum.in.ua;not null"`
	Desc  *string
	Title *string `gorm:"index"`
}

// TableName overrides the table name used by Article to `sum_articles`
func (Article) TableName() string {
	return "sum_articles"
}

type ArticleListFilter struct {
	crud.ListFilter
	Search string
}

type ImportLink struct {
	gorm.Model
	URL        string `gorm:"uniqueIndex"`
	Type       string
	HTML       *string
	Word       *string `gorm:"index"`
	ParentID   *uint
	ParentURL  *string
	Vocabulary string `gorm:"default:sum.in.ua;not null"`
	Desc       *string
	Title      *string `gorm:"index"`
}

// TableName overrides the table name used by ImportLink to `links`
func (ImportLink) TableName() string {
	return "links"
}
