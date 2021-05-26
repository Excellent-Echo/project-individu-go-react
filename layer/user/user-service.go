package user

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput) (UserFormat, error)
	GetUserByID(id string) (UserFormat, error)
	UpdateUserByID(id string, dataInput entity.UpdateUserInput) (UserFormat, error)
	DeleteByUserID(id string) (interface{}, error)
	LoginUser(input entity.LoginUserInput) (entity.User, error)
}

type userService struct {
	repository UserRepository
}

func UserNewService(repository UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) LoginUser(input entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %v not found", user.ID)
		return user, errors.New(newError)
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("invalid password")
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]UserFormat, error) {
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

func (s *userService) SaveNewUser(user entity.UserInput) (UserFormat, error) {
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

func (s *userService) GetUserByID(id string) (UserFormat, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return UserFormat{}, errors.New(newError)
	}

	userFormat := FormattingUser(user)

	return userFormat, nil

}

func (s *userService) UpdateUserByID(id string, dataInput entity.UpdateUserInput) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return UserFormat{}, errors.New(newError)
	}

	if dataInput.FirstName != "" || len(dataInput.FirstName) != 0 {
		dataUpdate["first_name"] = dataInput.FirstName
	}
	if dataInput.LastName != "" || len(dataInput.LastName) != 0 {
		dataUpdate["last_name"] = dataInput.LastName
	}
	if dataInput.UserName != "" || len(dataInput.UserName) != 0 {
		dataUpdate["user_name"] = dataInput.UserName
	}
	if dataInput.Email != "" || len(dataInput.Email) != 0 {
		dataUpdate["email"] = dataInput.Email
	}

	dataUpdate["updated_at"] = time.Now()

	fmt.Println(dataUpdate)

	userUpdated, err := s.repository.UpdateUser(id, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormattingUser(userUpdated)

	return formatUser, nil
}

func (s *userService) DeleteByUserID(id string) (interface{}, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return nil, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteUser(id)

	if err != nil {
		panic(err)
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("delete user id %s succeed", id)

	formatDelete := FormatDelete(msg)

	return formatDelete, nil
}
