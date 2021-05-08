package user

import (
	"projectpenyewaanlapangan/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
