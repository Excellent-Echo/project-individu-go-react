package user

import (
	"project-individu-go-react/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	GetOneUser(id string) (entity.User, error)
	// UpdateUser(id string, dataUpdate map[string]interface{}) (entity.User, error)
	// DeleteUser(id int) (entity.User, error)
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

func (r *repository) GetOneUser(id string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("user_id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// func (r *repository) UpdateUser(ID string, dataUpdate map[string]interface{}) (entity.User, error) {

// 	var user entity.User

// 	if err := r.db.Model(&user).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
// 		return user, err
// 	}

// 	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (r *repository) DeleteUser(id int) (entity.User, error) {
// 	var user entity.User

// 	r.db.Where("user_id = ?", id).Find(&user)

// 	if err := r.db.Where("user_id = ?", id).Delete(&entity.User{}).Error; err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }
