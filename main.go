package main

import (
	"project-individu-go-react/config"
	"project-individu-go-react/handler"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	userHandler             = handler.NewUserHandler(userService)
)

func main() {
	r := gin.Default()
	r.GET("/users", userHandler.ShowAllUser)
	// r.GET("/users", handler.GetAllUser)
	// r.GET("/users/:user_id", handler.HandleUsersID)
	r.POST("/users", handler.CreateNewUser)
	// r.DELETE("/users/:user_id", handler.HandleDeleteUser)
	// r.POST("/users/:user_id", handler.HandleUpdateUser)

	r.Run(":8080")
}
