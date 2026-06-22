package crud

import "gorm.io/gorm"

type CRUDModel interface {
	GetID() uint
}

type CRUDListFilter interface {
	GetOffset() int
	GetLimit() int
}

type CRUDRepository[M CRUDModel, F any] interface {
	GetDB() *gorm.DB
	Create(item M) (createdItem *M, err error)
	Update(id uint, item M) (*M, error)
	Get(id uint) (item *M, err error)
	GetMany(filter F) (items []M, err error)
	GetTotal(filter F) (count int64, err error)
	Delete(id uint) (err error)
}

type CRUDService[M CRUDModel, F any] interface {
	GetRepo() CRUDRepository[M, F]
	Create(item M) (createdItem *M, err error)
	Update(id uint, item M) (*M, error)
	Get(id uint) (item *M, err error)
	GetManyWithTotal(filter F) (items []M, total int64, err error)
	Delete(id uint) (err error)
}

type CRUDCreateRequest[M CRUDModel] interface {
	ToModel() M
}

type CRUDUpdateRequest[M CRUDModel] interface {
	ToModel() M
}

type CRUDListFilterRequest[F CRUDListFilter] interface {
	ToFilter() F
}
