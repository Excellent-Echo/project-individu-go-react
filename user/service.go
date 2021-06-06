package user

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	userdetail "project-individu-go-react/userDetail"
	"project-individu-go-react/userProfile"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUser() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput) (UserFormat, error)
	GetUserByID(id string) (UserDetailFormat, error)
	DeleteByUserID(id string) (interface{}, error)
	UpdateUserByID(id string, dataUpdate entity.UpdateUserInput) (UserFormat, error)
	LoginUser(userInput entity.LoginUserInput) (entity.User, error)
}

type service struct {
	repository      Repository
	userDetailRepo  userdetail.Repository
	userProfileRepo userProfile.Repository
}

func NewService(repo Repository, userDetailRepo userdetail.Repository, userProfileRepo userProfile.Repository) *service {
	return &service{repo, userDetailRepo, userProfileRepo}
}

func (s *service) GetAllUser() ([]UserFormat, error) {
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

//function service for create new user
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
		UpdateAt:  time.Now(),
	}

	createUser, err := s.repository.Create(newUser)

	formatUser := FormatUser(createUser)
	if err != nil {
		return formatUser, err
	}
	return formatUser, nil

}

//function service for get data user by id
func (s *service) GetUserByID(id string) (UserDetailFormat, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return UserDetailFormat{}, err
	}

	userDetail, err := s.userDetailRepo.FindByUserID(id)
	userProfile, err := s.userProfileRepo.FindByUserID(id)
	user, err := s.repository.FindByID(id)

	if err != nil {
		return UserDetailFormat{}, err
	}

	var userData = entity.UserDetailOutput{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdateAt,
		UserDetail:  userDetail,
		UserProfile: userProfile,
	}

	if user.ID == 0 {
		return UserDetailFormat{}, errors.New("user id not found")
	}
	formatUser := FormatDetailUser(userData)
	return formatUser, nil
}

//function service delete user by id
func (s *service) DeleteByUserID(id string) (interface{}, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return nil, err
	}

	user, err := s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("user id not found")
	}

	status, err := s.repository.DeleteByID(id)
	if err != nil {
		return nil, errors.New("error delete in internal server")
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", id)
	formatDelete := FormatDeleteUser(msg)

	return formatDelete, nil
}

func (s *service) UpdateUserByID(id string, dataInput entity.UpdateUserInput) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindByID(id)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		return UserFormat{}, errors.New("user id not found")
	}

	if dataInput.FirstName != "" || len(dataInput.FirstName) != 0 {
		dataUpdate["first_name"] = dataInput.FirstName
	}
	if dataInput.LastName != "" || len(dataInput.FirstName) != 0 {
		dataUpdate["last_name"] = dataInput.LastName
	}
	if dataInput.Email != "" || len(dataInput.Email) != 0 {
		dataUpdate["email"] = dataInput.Email
	}

	dataUpdate["update_at"] = time.Now()

	userUpdated, err := s.repository.UpdateByID(id, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormatUser(userUpdated)

	return formatUser, nil
}

func (s *service) LoginUser(userInput entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.FindByEmail(userInput.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	//checking password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		return user, errors.New("password invalid")
	}

	return user, nil
}
