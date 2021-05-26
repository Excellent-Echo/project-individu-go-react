package course

import (
	"project-individu-go-react/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Course, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Course, error) {
	var courses []entity.Course

	if err := r.db.Find(&courses).Error; err != nil {
		return courses, err
	}
	return courses, nil
}
