package user

import "gorm.io/gorm"

type Repository interface {
	InsertUser(user User) (User, error)
	FindUserByEmail(email string) (User, error)
	FindUserById(id int) (User, error)
	UpdateUser(user User) (User, error)
	FindUsers() (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindUserByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindUserById(id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindUsers() (User, error) {
	var users User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}