package category

import (
	"project-individu-go-react/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]entity.Categories, error)
	NewCategory(category entity.Categories) (entity.Categories, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAll() ([]entity.Categories, error) {
	var Category []entity.Categories

	err := r.db.Find(&Category).Error
	if err != nil {
		return Category, err
	}

	return Category, nil
}

func (r *Repository) NewCategory(category entity.Categories) (entity.Categories, error) {
	if err := r.db.Create(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}
