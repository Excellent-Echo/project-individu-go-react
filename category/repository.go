package category

import (
	"project-individu-go-react/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Category, error)
	FindByID(id string) (entity.Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Category, error) {
	var Categories []entity.Category

	if err := r.db.Find(&Categories).Error; err != nil {
		return Categories, err
	}
	return Categories, nil
}

func (r *repository) FindByID(id string) (entity.Category, error) {
	var category entity.Category

	if err := r.db.Where("id = ?", id).Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}
