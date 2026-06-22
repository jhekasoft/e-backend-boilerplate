package crud

type Service[M CRUDModel, F CRUDListFilter] struct {
	repo CRUDRepository[M, F]
}

func NewService[M CRUDModel, F CRUDListFilter](repo CRUDRepository[M, F]) *Service[M, F] {
	return &Service[M, F]{repo}
}

func (s *Service[M, F]) GetRepo() CRUDRepository[M, F] {
	return s.repo
}

func (s *Service[M, F]) Create(item M) (*M, error) {
	return s.repo.Create(item)
}

func (s *Service[M, F]) Update(id uint, item M) (*M, error) {
	return s.repo.Update(id, item)
}

func (s *Service[M, F]) Get(id uint) (*M, error) {
	return s.repo.Get(id)
}

func (s *Service[M, F]) GetManyWithTotal(filter F) (items []M, total int64, err error) {
	items, err = s.repo.GetMany(filter)
	if err != nil {
		return
	}

	total, err = s.repo.GetTotal(filter)
	return
}

func (s *Service[M, F]) Delete(id uint) (err error) {
	return s.repo.Delete(id)
}
