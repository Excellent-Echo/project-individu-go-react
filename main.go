package main

import (
	"project-individu-go-react/config"
	"project-individu-go-react/handler"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connect()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	userHandler             = handler.NewUserHandler(userService)
)

func main() {
	r := gin.Default()

	r.GET("/users", userHandler.ShowAllUsers)
	// r.POST("/users", handler.CreateNewUser)
	// r.GET("/users/:user_id", handler.GetUserByID)
	// r.PUT("/users/:user_id", handler.UpdateUserByID)
	// r.DELETE("/users/:user_id", handler.DeleteByUserID)

	r.Run(":4444")
}
