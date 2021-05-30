package userprofile

import (
	"gorm.io/gorm"
	"project-individu-go-react/entity"
)

type Repository interface {
	FindByUserProfileID(userID string) (entity.UserProfile, error)
	CreateUserProfile(userProfile entity.UserProfile) (entity.UserProfile, error)
	UpdateUserProfileByID(ID string, dataUserProfileUpdate map[string]interface{}) (entity.UserProfile, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUserProfileID(userID string) (entity.UserProfile, error) {
	var userProfile entity.UserProfile
	if err := r.db.Where("user_id = ?", userID).Find(&userProfile).Error; err != nil {
		return userProfile, err
	}
	return userProfile, nil
}

func (r *repository) CreateUserProfile(userProfile entity.UserProfile) (entity.UserProfile, error) {
	if err := r.db.Create(&userProfile).Error; err != nil {
		return userProfile, err
	}
	return userProfile, nil
}

func (r *repository) UpdateUserProfileByID(ID string, dataUserProfileUpdate map[string]interface{}) (entity.UserProfile, error) {
	var userProfile entity.UserProfile
	if err := r.db.Model(&userProfile).Where("id = ?", ID).Updates(dataUserProfileUpdate).Error; err != nil {
		return userProfile, err
	}
	if err := r.db.Where("id = ?", ID).Find(&userProfile).Error; err != nil {
		return userProfile, err
	}
	return userProfile, nil
}
