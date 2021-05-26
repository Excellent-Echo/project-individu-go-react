package handler

import (
	course "project-individu-go-react/courses"
	"project-individu-go-react/helper"

	"github.com/gin-gonic/gin"
)

type courseHandler struct {
	courseService course.Service
	// authService   auth.Service
}

func NewCourseHandler(courseService course.Service) *courseHandler {
	return &courseHandler{courseService}
}

func (h *courseHandler) ShowCourseHandler(c *gin.Context) {
	courses, err := h.courseService.GetAllCourse()

	if err != nil {
		responseErr := helper.APIResponse("internal server error", 500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}
	response := helper.APIResponse("success get all data course", 200, "status ok", courses)
	c.JSON(200, response)
}
