package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"project-go-react/entity"
	"project-go-react/helper"
	"time"
)

type Service interface {
	GetAllUser() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput) (UserFormat, error)
	GetUserByID(ID string) (UserFormat, error)
	DeleteByUserID(ID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllUser() ([]UserFormat, error) {
	//business logic
	users, err := s.repository.FindAll()
	var formatUsers []UserFormat

	for _, user := range users {
		formatUser := FormatUser(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}

func (s *service) SaveNewUser(user entity.UserInput) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	createUser, err := s.repository.Create(newUser)
	formatUser := FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil

}

func (s *service) GetUserByID(ID string) (UserFormat, error) {
	if err := helper.ValidateIDNumber(ID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindByID(ID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %d not found", user.ID)

		return UserFormat{}, errors.New(newError)
	}

	formatUser := FormatUser(user)
	return formatUser, nil
}

func (s *service) DeleteByUserID(ID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(ID); err != nil {
		return nil, err
	}

	user, err := s.repository.FindByID(ID)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %d not found", user.ID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(ID)

	if status == "error" {
		return nil, errors.New("error delete")
	}

	msg := fmt.Sprintf("succes delete user ID : %s", ID)

	formatDelete := FormatDeleteUser(msg)

	return formatDelete, nil

}
