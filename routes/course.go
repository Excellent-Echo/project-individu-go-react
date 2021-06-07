package routes

import (
	course "project-individu-go-react/courses"
	"project-individu-go-react/handler"

	"github.com/gin-gonic/gin"
)

var (
	courseRepository = course.NewRepository(DB)
	courseService    = course.NewService(courseRepository)
	courseHandler    = handler.NewCourseHandler(courseService)
)

func CourseRoute(r *gin.Engine) {
	r.GET("/courses", courseHandler.ShowCourseHandler)
	r.GET("/courses/:id", courseHandler.ShowCourseByIDHandler)
}
