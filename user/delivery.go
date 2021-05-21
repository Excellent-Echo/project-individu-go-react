package user

import (
	"project-individu-go-react/helper"

	"github.com/gin-gonic/gin"
)

type userDeliver struct {
	userService Service
}

func NewUserDeliver(userService Service) *userDeliver {
	return &userDeliver{userService}
}

func (d *userDeliver) ShowUserDeliver(c *gin.Context) {
	users, err := d.userService.GetAllUser()

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responError)

		return
	}

	response := helper.APIResponse("success get all user", 200, "status OK", users)
	c.JSON(200, response)
}
