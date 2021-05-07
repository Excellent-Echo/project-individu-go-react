package handler

import (
	"net/http"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) ShowAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in internal server",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
