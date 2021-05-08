package user

import (
	"gorm.io/gorm"
	"project-individu-go-react/entity"
)

type Repository interface {
	FindAllUsers() ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	FindUserByID(ID string) (entity.User, error)
	FindandDeleteUserByID(ID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// FindAllUsers fungsi untuk mengambil semua data di model users menggunakan find
func (r *repository) FindAllUsers() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// CreateUser fungsi untuk membuat data baru di model users menggunakan create
func (r *repository) CreateUser(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, nil
	}
	return user, nil
}

// FindUserByID fungsi untuk mencari data user berdasarkan "id" di model users menggunakan Where dan Find
func (r *repository) FindUserByID(ID string) (entity.User, error) {
	var userFind entity.User
	if err := r.db.Where("id = ?", ID).Find(&userFind).Error; err != nil {
		return userFind, err
	}
	return userFind, nil
}

// FindandDeleteUserByID fungsi untuk menghapus data user berdasarkan "id" di model users menggunakan Where dan Delete
func (r *repository) FindandDeleteUserByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.User{}).Error; err != nil {
		return "error", err
	}
	return "success", nil
}
