package handler

import (
	"net/http"
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
