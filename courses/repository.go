package course

import (
	"project-individu-go-react/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Course, error)
	FindByID(id string) (entity.Course, error)
	FindByCategoryID(id string) ([]entity.Course, error)
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

func (r *repository) FindByID(id string) (entity.Course, error) {
	var course entity.Course

	if err := r.db.Where("id = ?", id).Find(&course).Error; err != nil {
		return course, err
	}
	return course, nil
}

func (r *repository) FindByCategoryID(id string) ([]entity.Course, error) {
	var courses []entity.Course

	if err := r.db.Where("category_id = ?", id).Find(&courses).Error; err != nil {
		return courses, err
	}
	return courses, nil
}
