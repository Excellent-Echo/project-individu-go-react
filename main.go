package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"project-go-react/config"
	"project-go-react/handler"
	"project-go-react/user"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	userHandler             = handler.NewUserHandler(userService)
)

func main() {
	r := gin.Default()

	r.GET("/users", userHandler.ShowUserHandler)
	r.GET("/users/:user_id", userHandler.GetUserByIDHandler)
	r.POST("/users", userHandler.CreateUserHandler)

	// r.POST("/users", handler.CreateNewUser)
	// r.PUT("/users/:user_id", handler.UpdateByUserID)
	r.DELETE("/users/:user_id", userHandler.DeleteUserByIDHandler)

	r.Run(":1212")

}
