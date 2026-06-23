package models

import (
	"time"

	"gorm.io/gorm"

	internalModels "e-backend-boilerplate/pkg/ebackend/models"
)

type User struct {
	gorm.Model
	Username    string     `gorm:"uniqueIndex"`
	Email       string     `gorm:"uniqueIndex"`
	Name        string     `gorm:"index"`
	ActivatedAt *time.Time `gorm:"index"`
	Password    string     `json:"-"`
}

type AuthClaims internalModels.AuthClaims
