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

func (h *courseHandler) ShowCourseByIDHandler(c *gin.Context) {
	id := c.Param("id")

	user, err := h.courseService.GetCourseByID(id)
	if err != nil {
		responseErr := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseErr)
		return
	}
	response := helper.APIResponse("success get data user by id", 200, "success", user)
	c.JSON(200, response)
}
