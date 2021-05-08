package user

import (
	"project-individu-go-react/entity"
	"time"
)

type Service interface {
	GetAllUser() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput) (UserFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUser() ([]UserFormat, error) {
	users, err := s.repository.GetAll()
	var formatUsers []UserFormat

	for _, user := range users {
		formatuser := FormatUser(user)
		formatUsers = append(formatUsers, formatuser)
	}

	if err != nil {
		return formatUsers, err
	}
	return formatUsers, nil
}

func (s *service) SaveNewUser(user entity.UserInput) (UserFormat, error) {

	var newUser = entity.User{
		RoleID:    user.RoleID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.repository.CreateUser(newUser)
	formatUser := FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}
