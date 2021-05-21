package user

import (
	"project-individu-go-react/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.User, error)
	Create(user entity.User) (entity.User, error)
	FindByID(id string) (entity.User, error)
	DeleteByID(id string) (string, error)
	UpdateByID(id string, dataUpdate map[string]interface{}) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// function repository for Get All data User
func (r *repository) FindAll() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

//function repository for create new users
func (r *repository) Create(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//function repository for get data user by id
func (r *repository) FindByID(id string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//function repository for delete user by id
func (r *repository) DeleteByID(id string) (string, error) {
	if err := r.db.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return "error", err
	}
	return "success", nil
}

func (r *repository) UpdateByID(id string, dataUpdate map[string]interface{}) (entity.User, error) {
	var user entity.User

	if err := r.db.Model(&user).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
