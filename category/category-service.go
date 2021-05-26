package category

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
)

type CategoryService interface {
	GetAllCategories() ([]entity.Categories, error)
	SaveNewCategory(category entity.CategoryInput) (entity.Categories, error)
	FindCategoryByName(categoryName string) (entity.Categories, error)
}

type categoryService struct {
	repository CategoryRepository
}

func NewService(repository CategoryRepository) *categoryService {
	return &categoryService{repository}
}

func (s *categoryService) GetAllCategories() ([]entity.Categories, error) {
	categories, err := s.repository.GetAll()

	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (s *categoryService) SaveNewCategory(category entity.CategoryInput) (entity.Categories, error) {
	var newCategory = entity.Categories{
		CategoryName: category.CategoryName,
	}

	createCategory, err := s.repository.NewCategory(newCategory)

	if err != nil {
		return createCategory, err
	}

	return createCategory, nil
}

func (s *categoryService) FindCategoryByName(categoryName string) (entity.Categories, error) {
	category, err := s.repository.FindCategoryName(categoryName)

	if err != nil {
		return entity.Categories{}, err
	}

	if category.CategoryName == "" {
		newError := fmt.Sprintf("category %s is not found", categoryName)
		return entity.Categories{}, errors.New(newError)
	}

	return category, nil
}
