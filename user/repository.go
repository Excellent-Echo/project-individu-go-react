package user

import (
	"project-individu-go-react/entity"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	GetOneUser(id int) (entity.User, error)
	UpdateUser(id int) (entity.User, error)
	DeleteUser(id int) (entity.User, error)
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

func (r *repository) CreateUser(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetOneUser(id int) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("user_id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateUser(id int) (entity.User, error) {
	var user entity.User
	var userInput entity.UserInput

	if err := r.db.Where("user_id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	user.FirstName = userInput.FirstName
	user.LastName = userInput.LastName
	user.UserName = userInput.UserName
	user.Email = userInput.Email
	user.Password = userInput.Password
	user.UpdatedAt = time.Now()

	if err := r.db.Where("user_id = ?", id).Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteUser(id int) (entity.User, error) {
	var user entity.User

	r.db.Where("user_id = ?", id).Find(&user)

	if err := r.db.Where("user_id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return user, err
	}

	return user, nil
}
