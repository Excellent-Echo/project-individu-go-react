package userProfile

import (
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
)

type Service interface {
	GetUserProfile(userID string) (entity.UserProfile, error)
	SaveNewUserProfile(pathFile string, userID int) (entity.UserProfile, error)
	UpdateUserProfileByID(pathFile string, userProfileID string) (entity.UserProfile, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserProfile(userID string) (entity.UserProfile, error) {
	userProfile, err := s.repository.FindByUserID(userID)

	if err != nil {
		return userProfile, err
	}

	return userProfile, nil
}

func (s *service) SaveNewUserProfile(pathFile string, userID int) (entity.UserProfile, error) {
	var newUserProfile = entity.UserProfile{
		ProfileUser: pathFile,
		UserID:      userID,
	}

	userProfile, err := s.repository.Create(newUserProfile)

	if err != nil {
		return userProfile, err
	}

	return userProfile, nil

}

func (s *service) UpdateUserProfileByID(pathFile string, userProfileID string) (entity.UserProfile, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(userProfileID); err != nil {
		return entity.UserProfile{}, err
	}

	dataUpdate["profile_user"] = pathFile

	userProfileUpdate, err := s.repository.UpdateByUserID(userProfileID, dataUpdate)

	if err != nil {
		return userProfileUpdate, err
	}
	return userProfileUpdate, nil

}
