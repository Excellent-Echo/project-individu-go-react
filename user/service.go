package user

import (
	"project-individu-go-react/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUsers() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput) (UserFormat, error)
	GetUserByID(id string) (UserFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUsers() ([]UserFormat, error) {
	users, err := s.repository.GetAll()

	var usersFormat []UserFormat

	for _, user := range users {
		var userFormat = FormattingUser(user)
		usersFormat = append(usersFormat, userFormat)
	}

	if err != nil {
		return usersFormat, err
	}

	return usersFormat, nil
}

func (s *service) SaveNewUser(user entity.UserInput) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  string(genPassword),
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

func (s *service) GetUserByID(id string) (UserFormat, error) {
	user, err := s.repository.GetOneUser(id)

	userFormat := FormattingUser(user)

	if err != nil {
		return userFormat, err
	}

	return userFormat, nil

}
