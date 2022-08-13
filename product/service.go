package product

type Service interface {
	FetchProducts()
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FetchProducts() {
	s.repository.FetchProduct()
}
