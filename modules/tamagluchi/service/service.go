package service

import (
	"e-backend-boilerplate/modules/tamagluchi/models"
	"e-backend-boilerplate/modules/tamagluchi/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo}
}

func (s *Service) CreatePet(pet models.Pet) (models.TamagluchiState, error) {
	newPet := models.Pet{
		Name: pet.Name,
		Age:  0,
		Type: pet.Type,
	}

	state := models.TamagluchiState{
		Pet: newPet,
		Main: models.PetMainValues{
			Food:   40,
			Water:  30,
			Rest:   85,
			Joy:    90,
			Health: 95,
		},
		Secondary: models.PetSecondaryValues{
			IsResting: false,
		},
		House: models.HouseValues{
			IsHeaped: false,
		},
	}

	return state, nil
}

func (s *Service) Calculate(state models.TamagluchiState, period int) (models.TamagluchiState, error) {
	// TODO: implement calculation logic
	return state, nil
}
