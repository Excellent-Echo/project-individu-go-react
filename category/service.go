package category

import (
	"errors"
	"fmt"
	course "project-individu-go-react/courses"
	"project-individu-go-react/entity"
)

type Service interface {
	GetAllCategory() ([]entity.Category, error)
	GetCategoryByID(id string) (CategoryDetailFormat, error)
}

type service struct {
	repository       Repository
	courseRepository course.Repository
}

func NewService(repo Repository, course course.Repository) *service {
	return &service{repo, course}
}

func (s *service) GetAllCategory() ([]entity.Category, error) {
	// Find all data hire from database
	Categories, err := s.repository.FindAll()

	if err != nil {
		return Categories, err
	}

	return Categories, nil
}

func (s *service) GetCategoryByID(id string) (CategoryDetailFormat, error) {

	category, err := s.repository.FindByID(id)
	course, err := s.courseRepository.FindByCategoryID(id)

	if err != nil {
		return CategoryDetailFormat{}, err
	}

	var categoryData = CategoryDetailFormat{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Courses:     course,
	}

	if category.ID == 0 {
		newError := fmt.Sprintf("category id %s not found", id)
		return CategoryDetailFormat{}, errors.New(newError)
	}

	return categoryData, nil
}
