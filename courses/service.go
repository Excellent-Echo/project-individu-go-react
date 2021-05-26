package course

type Service interface {
	GetAllCourse() ([]CourseFormat, error)
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
