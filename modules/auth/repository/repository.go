package repository

import (
	"e-backend-boilerplate/modules/auth/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(item models.User) (createdItem *models.User, err error) {
	if err := r.db.Create(&item).Error; err != nil {
		return nil, err
	}

	createdItem, err = r.Get(item.ID)
	return
}

func (r *Repository) Update(id uint, item models.User) (*models.User, error) {
	var updatedItem models.User
	if err := r.db.Where("id = ?", id).Updates(&item).Scan(&updatedItem).Error; err != nil {
		return nil, err
	}

	return &updatedItem, nil
}

func (r *Repository) Get(id uint) (item *models.User, err error) {
	err = r.db.First(&item, id).Error
	return
}

func (r *Repository) Delete(id uint) (err error) {
	err = r.db.Delete(&models.User{}, id).Error
	return
}

func (r *Repository) FindByUsernameOrEmail(credential string) (item *models.User, err error) {
	err = r.db.
		Where("username = ?", credential).
		Or("email = ?", credential).
		First(&item).
		Error
	return
}
