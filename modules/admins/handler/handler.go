package handler

import (
	"e-backend-boilerplate/modules/admins/models"

	"github.com/jhekasoft/e-backend/crud"
)

type Handler struct {
	crud.Handler[
		models.Admin,
		models.AdminListFilter,
		CreateAdminRequest,
		UpdateAdminRequest,
		AdminListFilter,
	]
}

func NewHandler(service crud.CRUDService[models.Admin, models.AdminListFilter]) *Handler {
	return &Handler{
		*crud.NewHandler[
			models.Admin,
			models.AdminListFilter,
			CreateAdminRequest,
			UpdateAdminRequest,
			AdminListFilter,
		](service),
	}
}

type AdminListFilter struct {
	Offset int               `query:"Offset"`
	Limit  int               `query:"Limit"`
	Role   *models.AdminRole `query:"Role"`
	Search string            `query:"Search"`
}

func (req AdminListFilter) ToFilter() models.AdminListFilter {
	return models.AdminListFilter{
		ListFilter: crud.ListFilter{
			Offset: req.Offset,
			Limit:  req.Limit,
		},
		Role:   req.Role,
		Search: req.Search,
	}
}

type CreateAdminRequest struct {
	Username string `validate:"required"`
	Name     string
	Role     models.AdminRole `validate:"required"`
	Password string           `validate:"required,gte=6"`
}

func (req CreateAdminRequest) ToModel() models.Admin {
	return models.Admin{
		Username: req.Username,
		Name:     req.Name,
		Role:     req.Role,
		Password: req.Password,
	}
}

type UpdateAdminRequest struct {
	Name string           `validate:"required"`
	Role models.AdminRole `validate:"required"`
}

func (req UpdateAdminRequest) ToModel() models.Admin {
	return models.Admin{
		Name: req.Name,
		Role: req.Role,
	}
}
