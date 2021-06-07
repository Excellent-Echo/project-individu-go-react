package course

import "project-individu-go-react/entity"

type CourseFormat struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func FormatCourse(course entity.Course) CourseFormat {
	var formatCourse = CourseFormat{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
	}
	return formatCourse
}
