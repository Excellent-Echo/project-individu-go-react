package user

import (
	"project-go-react/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]entity.User, error) {
	var Users []entity.User
	if err := r.db.Find(&Users).Error; err != nil {
		return Users, err
	}
	return Users, nil
}
