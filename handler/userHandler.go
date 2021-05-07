package handler

import (
	"net/http"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) ShowAllUsersHandler(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in internal server",
		})
		return
	}

	userResponse := helper.APIResponse("get all users succeed", 200, "success", users)
	c.JSON(http.StatusOK, userResponse)
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "error input data",
		})
		return
	}

	response, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "error input data",
		})
		return
	}

	userResponse := helper.APIResponse("insert user succeed", 201, "success", response)
	c.JSON(http.StatusOK, userResponse)
}
