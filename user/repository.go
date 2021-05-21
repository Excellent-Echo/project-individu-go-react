package user

import (
	"project-individu-go-react/entities"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entities.User, error)
	//Create(user entities.User) (entities.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entities.User, error) {
	var users []entities.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
