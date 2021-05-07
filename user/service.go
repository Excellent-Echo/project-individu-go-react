package user

import (
	"project-individu-go-react/entity"
	"time"
)

type Service interface {
	GetAllUsers() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput) (UserFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUsers() ([]UserFormat, error) {
	Users, err := s.repository.GetAll()

	var usersFormat []UserFormat

	for _, user := range Users {
		var userFormat = FormattingUser(user)
		usersFormat = append(usersFormat, userFormat)
	}

	if err != nil {
		return usersFormat, err
	}

	return usersFormat, nil
}

func (s *service) SaveNewUser(user entity.UserInput) (UserFormat, error) {
	var newUser = entity.User{
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.repository.CreateUser(newUser)
	formatUser := FormattingUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}
