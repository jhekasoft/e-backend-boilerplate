package models

import "e-backend/pkg/ebackend/crud"

// AdminRole is role of the administrator.
type AdminRole string

const (
	AdminRoleSuper   AdminRole = "super"
	AdminRoleDefault AdminRole = "default"
)

type Admin struct {
	crud.Model
	Username string    `gorm:"uniqueIndex"`
	Name     string    `gorm:"index"`
	Role     AdminRole `gorm:"index"`
	Password string    `json:"-"`
}

type AdminListFilter struct {
	crud.ListFilter
	Role   *AdminRole
	Search string
}
