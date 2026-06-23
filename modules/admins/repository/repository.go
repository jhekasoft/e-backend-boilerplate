package repository

import (
	"e-backend-boilerplate/modules/admins/models"
	"e-backend-boilerplate/pkg/ebackend/crud"

	"gorm.io/gorm"
)

type Repository struct {
	crud.Repository[models.Admin, models.AdminListFilter]
}

func NewRepository(db *gorm.DB) *Repository {
	listOrder := "id desc"
	var listScope crud.ListScopeFunc[models.AdminListFilter] = listScope
	return &Repository{*crud.NewRepository[models.Admin](db, &listScope, listOrder)}
}

func listScope(filter models.AdminListFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(filter.Search) > 0 {
			db.Where(
				db.Session(&gorm.Session{}).Unscoped().
					Where("name ILIKE ?", "%"+filter.Search+"%").
					Or("username ILIKE ?", "%"+filter.Search+"%"),
			)
		}
		if filter.Role != nil {
			db.Where("role = ?", filter.Role)
		}
		return db
	}
}
