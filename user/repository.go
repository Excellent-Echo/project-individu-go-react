package user

import (
	"project-individu-go-react/entity"

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

	err := r.db.Find(&Users).Error
	if err != nil {
		return Users, err
	}

	return Users, nil
}
