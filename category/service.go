package category

type Service interface {
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}
