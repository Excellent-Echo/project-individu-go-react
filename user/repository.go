package user

import (
	"project-individu-go-react/entities"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entities.User, error)
	Create(user entities.User) (entities.User, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entities.User, error)
	DeleteByID(ID string) (string, error)
	FindByID(ID string) (entities.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//USERSS =======
//Read USER
func (r *repository) FindAll() ([]entities.User, error) {
	var users []entities.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

//Read by ID
func (r *repository) FindByID(ID string) (entities.User, error) {
	var user entities.User

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

//Create
func (r *repository) Create(user entities.User) (entities.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

//Update
func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entities.User, error) {
	var user entities.User

	if err := r.db.Model(&user).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

//Delete
func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id=?", ID).Delete(&entities.User{}).Error; err != nil {
		return "error", err
	}

	return "delete success", nil
}
