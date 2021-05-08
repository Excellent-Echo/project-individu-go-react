package handler

import (
	"net/http"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
	"projectpenyewaanlapangan/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

//showUserHandler for handling show all user in db from route "/users"
func (h *userHandler) ShowUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`message`: `error in internal server error`,
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUserHandler for handing if user / external create new user from route "/users"
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new User", 201, "status Created", newUser)
	c.JSON(201, response)
}
