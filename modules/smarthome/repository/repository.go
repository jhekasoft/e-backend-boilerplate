package repository

import (
	"e-backend-boilerplate/modules/smarthome/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(item models.SmartHomeSensorValue) (createdItem *models.SmartHomeSensorValue, err error) {
	if err := r.db.Create(&item).Error; err != nil {
		return nil, err
	}

	createdItem, err = r.Get(item.ID)
	return
}

func (r *Repository) Get(id uint) (item *models.SmartHomeSensorValue, err error) {
	err = r.db.First(&item, id).Error
	return
}
