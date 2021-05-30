package userprofile

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"strconv"
)

type Service interface {
	GetUserProfileByID(userID string) (entity.UserProfile, error)
	MakeNewUserProfile(pathFile string, userID int) (entity.UserProfile, error)
	UpdateUserProfileByID(pathFile string, userID string) (entity.UserProfile, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserProfileByID(userID string) (entity.UserProfile, error) {
	userProfile, err := s.repository.FindByUserProfileID(userID)
	if err != nil {
		return userProfile, err
	}
	return userProfile, nil
}

func (s *service) MakeNewUserProfile(pathFile string, userID int) (entity.UserProfile, error) {
	ID := strconv.Itoa(userID)
	checkingUserProfile, _ := s.repository.FindByUserProfileID(ID)
	if checkingUserProfile.UserID == userID {
		StatusError := fmt.Sprintf("user profile for user id %d succesful created", userID)
		return checkingUserProfile, errors.New(StatusError)
	}

	addNewUserProfile := entity.UserProfile{
		ImageUser: pathFile,
		UserID:    userID,
	}
	userProfile, err := s.repository.CreateUserProfile(addNewUserProfile)
	if err != nil {
		return userProfile, err
	}
	return userProfile, nil
}

func (s *service) UpdateUserProfileByID(pathFile string, userID string) (entity.UserProfile, error) {
	var dataImageUpdate = map[string]interface{}{}
	if err := helper.ValidateIDNumber(userID); err != nil {
		return entity.UserProfile{}, err
	}
	dataImageUpdate["image_user"] = pathFile

	userProfileUpdate, err := s.repository.UpdateUserProfileByID(userID, dataImageUpdate)
	if err != nil {
		return userProfileUpdate, err
	}
	return userProfileUpdate, nil
}
