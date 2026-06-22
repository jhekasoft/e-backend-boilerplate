package crud

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
}

func (m Model) GetID() uint {
	return m.ID
}

type ListFilter struct {
	Offset int
	Limit  int
}

func (f ListFilter) GetOffset() int {
	return f.Offset
}

func (f ListFilter) GetLimit() int {
	return f.Limit
}

type ListScopeFunc[F CRUDListFilter] func(F) func(db *gorm.DB) *gorm.DB

type Repository[M CRUDModel, F CRUDListFilter] struct {
	db        *gorm.DB
	listScope *ListScopeFunc[F]
	listOrder any
}

func NewRepository[M CRUDModel, F CRUDListFilter](
	db *gorm.DB,
	listScope *ListScopeFunc[F],
	listOrder any,
) *Repository[M, F] {
	return &Repository[M, F]{db, listScope, listOrder}
}

func (r *Repository[M, F]) GetDB() *gorm.DB {
	return r.db
}

func (r *Repository[M, F]) Create(item M) (createdItem *M, err error) {
	if err := r.db.Create(&item).Error; err != nil {
		return nil, err
	}

	createdItem, err = r.Get(item.GetID())
	return
}

func (r *Repository[M, F]) Update(id uint, item M) (*M, error) {
	var updatedItem M
	if err := r.db.Where("id = ?", id).Updates(&item).Scan(&updatedItem).Error; err != nil {
		return nil, err
	}

	return &updatedItem, nil
}

func (r *Repository[M, F]) Get(id uint) (item *M, err error) {
	err = r.db.First(&item, id).Error
	return
}

func (r *Repository[M, F]) Delete(id uint) (err error) {
	var item M
	err = r.db.Delete(&item, id).Error
	return
}

func (r *Repository[M, F]) Paginate(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

func (r *Repository[M, F]) GetMany(filter F) (items []M, err error) {
	err = r.getListQuery(filter).
		Scopes(r.Paginate(filter.GetOffset(), filter.GetLimit())).
		Find(&items).
		Error
	return
}

func (r *Repository[M, F]) GetTotal(filter F) (count int64, err error) {
	err = r.getListQuery(filter).Count(&count).Error
	return
}

func (r *Repository[M, F]) getListQuery(filter F) *gorm.DB {
	var item M
	tx := r.GetDB().Model(&item)

	if r.listOrder != nil {
		tx.Order(r.listOrder)
	}

	if r.listScope != nil {
		listScopeFunc := *r.listScope
		tx.Scopes(listScopeFunc(filter))
	}

	return tx
}
