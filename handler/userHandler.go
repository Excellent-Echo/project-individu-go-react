package handler

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHander(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) ShowUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error internal server",
			"error":   err.Error(),
		})
		return
	}

	response := helper.APIResponse(
		"Success Get All Users Data",
		200,
		"Status OK",
		users,
	)

	c.JSON(200, response)
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var NewUserInput entity.UserInput

	if err := c.ShouldBindJSON(&NewUserInput); err != nil {
		splitError := helper.SplitErrorInformation(err)

		responseError := helper.APIResponse(
			"Input data required",
			400,
			"bad request",
			gin.H{
				"error": splitError,
			},
		)

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(NewUserInput)

	if err != nil {
		responseError := helper.APIResponse(
			"Internal server error",
			500,
			"error",
			gin.H{
				"error": err.Error(),
			},
		)

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse(
		"Success Create new Users Data",
		201,
		"Status Created",
		newUser,
	)

	c.JSON(201, response)
}
