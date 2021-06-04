package handler

import (
	"net/http"
	"project-go-react/helper"
	"project-go-react/user"

	"github.com/gin-gonic/gin"
)

type userhandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userhandler {
	return &userhandler{userService}
}

func (h *userhandler) ShowAllUser(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in internal server",
		})
		return
	}
	userResponse := helper.APIResponse("success get all user", 200, "success", users)
	c.JSON(http.StatusOK, userResponse)
}
