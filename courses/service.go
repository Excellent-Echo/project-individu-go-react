package course

import (
	"errors"
	"project-individu-go-react/helper"
)

type Service interface {
	GetAllCourse() ([]CourseFormat, error)
	GetCourseByID(id string) (CourseFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllCourse() ([]CourseFormat, error) {
	courses, err := s.repository.FindAll()

	var formatCourses []CourseFormat

	for _, course := range courses {
		formatCourse := FormatCourse(course)
		formatCourses = append(formatCourses, formatCourse)
	}
	if err != nil {
		return formatCourses, err
	}
	return formatCourses, nil
}

func (s *service) GetCourseByID(id string) (CourseFormat, error) {
	// validate input ID is not negative number
	if err := helper.ValidateIDNumber(id); err != nil {
		return CourseFormat{}, err
	}

	course, err := s.repository.FindByID(id)

	if err != nil {
		return CourseFormat{}, err
	}

	if course.ID == 0 {
		return CourseFormat{}, errors.New("job course not found")
	}
	formatCourse := FormatCourse(course)
	return formatCourse, nil
}
