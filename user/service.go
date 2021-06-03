package user

import "project-go-react/entity"

type Service interface {
	GetAllUser() ([]entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUser() ([]entity.User, error) {
	Users, err := s.repository.GetAll()

	if err != nil {
		return Users, err
	}

	return Users, nil
}
