package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"time"
)

type Service interface {
	GetAllUser() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput) (UserFormat, error)
	GetUserByID(userID string) (UserFormat, error)
	GetandDeleteUserByID(userID string) (interface{}, error)
	GetandUpdateUserByID(userID string, dataUserInput entity.UserInputUpdate) (UserFormat, error)
	LoginUserbyEmail(input entity.UserLoginInput) (entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// GetAllUser business logic untuk mengambil semua data di model users
func (s *service) GetAllUser() ([]UserFormat, error) {
	users, err := s.repository.FindAllUsers()
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

// SaveNewUser business logic untuk membuat data baru ke model users
func (s *service) SaveNewUser(user entity.UserInput) (UserFormat, error) {

	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		RoleID:    user.RoleID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  string(genPassword),
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

// GetUserByID business logic untuk mencari data user berdasarkan "id: dari model users
func (s *service) GetUserByID(userID string) (UserFormat, error) {
	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindUserByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	formatUser := FormatUser(user)

	return formatUser, nil
}

// GetandDeleteUserByID business logic untuk menghapus data user berdasarkan "id" dari model users
func (s *service) GetandDeleteUserByID(userID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindUserByID(userID)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.FindandDeleteUserByID(userID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	message := fmt.Sprintf("success delete user by ID : %s", userID)

	formatDelete := FormatDeleteUser(message)

	return formatDelete, nil
}

// GetandUpdateUserByID business logic untuk mengupdate data user berdasarkan "id" dari model users
func (s *service) GetandUpdateUserByID(userID string, dataUserInput entity.UserInputUpdate) (UserFormat, error) {

	var dataUserUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindUserByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id not found : %s", userID)
		return UserFormat{}, errors.New(newError)
	}

	if dataUserInput.RoleID != 0 {
		dataUserUpdate["role_id"] = dataUserInput.RoleID
	}
	if dataUserInput.FirstName != "" || len(dataUserInput.FirstName) != 0 {
		dataUserUpdate["firstname"] = dataUserInput.FirstName
	}
	if dataUserInput.Lastname != "" || len(dataUserInput.Lastname) != 0 {
		dataUserUpdate["lastname"] = dataUserInput.Lastname
	}
	if dataUserInput.Email != "" || len(dataUserInput.Email) != 0 {
		dataUserUpdate["email"] = dataUserInput.Email
	}
	dataUserUpdate["updated_at"] = time.Now()

	fmt.Println(dataUserUpdate)

	userUpdated, err := s.repository.FindAndUpdateUserByID(userID, dataUserUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormatUser(userUpdated)

	return formatUser, nil
}

func (s *service) LoginUserbyEmail(input entity.UserLoginInput) (entity.User, error) {
	user, err := s.repository.FindUserByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %v not found", user.ID)
		return user, errors.New(newError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("password invalid")
	}

	return user, nil
}
