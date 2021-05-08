package user

import (
	"gorm.io/gorm"
	"project-individu-go-react/entity"
)

type Repository interface {
	GetAll() ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) CreateUser(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, nil
	}
	return user, nil
}
