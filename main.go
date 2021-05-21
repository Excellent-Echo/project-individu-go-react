package main

import (
	"project-individu-go-react/config"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connect()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	userDelivery            = user.NewUserDeliver(userService)
)

func main() {
	r := gin.Default()

	r.GET("/users", userDelivery.ShowUserDeliver)
	r.GET("/users/:user_id", userDelivery.GetUserByIDDeliver)
	r.POST("users/register", userDelivery.CreateUserDeliver)
	r.PUT("users/:user_id", userDelivery.UpdateUserByIDDeliver)
	r.DELETE("users/:user_id", userDelivery.DeleteUserByIDDeliver)

	r.Run(":8888")
}
