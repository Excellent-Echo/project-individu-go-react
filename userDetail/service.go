package userdetail

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"strconv"
	"time"
)

type Service interface {
	GetUserDetailByUserID(userID string) (entity.UserDetail, error)
	SaveNewUserDetail(input entity.UserDetailInput, userID string) (entity.UserDetail, error)
	UpdateUserDetailByID(ID string, dataInput entity.UpdateUserDetailInput) (entity.UserDetail, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetUserDetailByUserID(userID string) (entity.UserDetail, error) {
	userDetail, err := s.repository.FindByUserID(userID)

	if err != nil {
		return userDetail, err
	}

	if userDetail.ID == 0 {
		errStatus := fmt.Sprintf("userdetail for user id %s not created", userID)
		return userDetail, errors.New(errStatus)
	}

	return userDetail, nil
}

func (s *service) SaveNewUserDetail(input entity.UserDetailInput, userID string) (entity.UserDetail, error) {
	IDUser, _ := strconv.Atoi(userID)

	checkStatus, err := s.repository.FindByID(userID)

	if err != nil {
		return checkStatus, err
	}

	if checkStatus.UserID == IDUser {
		errorStatus := fmt.Sprintf("userDetail for user ID : %s has been created", userID)
		return checkStatus, errors.New(errorStatus)
	}

	var NewUserDetail = entity.UserDetail{
		Address:     input.Address,
		NoHandphone: input.NoHandphone,
		Gender:      input.Gender,
		UserID:      IDUser,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createUserDetail, err := s.repository.Create(NewUserDetail)

	if err != nil {
		return createUserDetail, err
	}

	return createUserDetail, nil
}

func (s *service) UpdateUserDetailByID(ID string, dataInput entity.UpdateUserDetailInput) (entity.UserDetail, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(ID); err != nil {
		return entity.UserDetail{}, err
	}

	userDetail, err := s.repository.FindByID(ID)

	if err != nil {
		return entity.UserDetail{}, err
	}

	if userDetail.ID == 0 {
		newError := fmt.Sprintf("user Detail id %s not found", ID)
		return entity.UserDetail{}, errors.New(newError)
	}

	if dataInput.Address != "" || len(dataInput.Gender) != 0 {
		dataUpdate["address"] = dataInput.Address
	}
	if dataInput.NoHandphone != 0 {
		dataUpdate["no_handphone"] = dataInput.NoHandphone
	}
	if dataInput.Gender != "" || len(dataInput.Gender) != 0 {
		dataUpdate["dender"] = dataInput.Gender
	}

	dataUpdate["updated_at"] = time.Now()

	userDetailUpdate, err := s.repository.UpdateByID(ID, dataUpdate)

	if err != nil {
		return userDetailUpdate, err
	}

	return userDetailUpdate, nil
}
