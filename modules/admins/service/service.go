package service

import (
	"e-backend-boilerplate/modules/admins/models"
	"e-backend-boilerplate/pkg/ebackend/crud"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	crud.Service[models.Admin, models.AdminListFilter]
}

func NewService(repo crud.CRUDRepository[models.Admin, models.AdminListFilter]) *Service {
	return &Service{*crud.NewService(repo)}
}

func (s *Service) Create(item models.Admin) (createdItem *models.Admin, err error) {
	// Hash password
	passwordHash, err := s.hashPassword(item.Password)
	if err != nil {
		return
	}
	item.Password = passwordHash

	createdItem, err = s.GetRepo().Create(item)
	if err != nil {
		return
	}

	return
}

func (s *Service) hashPassword(password string) (hash string, err error) {
	hashBytes, err :=
		bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hash = string(hashBytes)
	return
}
