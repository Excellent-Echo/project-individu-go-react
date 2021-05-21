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

	r.Run(":8888")
}
