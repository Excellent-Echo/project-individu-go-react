package main

import (
	"project-individu-go-react/config"
	"project-individu-go-react/handler"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
)

var (
	DB             = config.Connection()
	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)
	userHandler    = handler.NewUserHandler(userService)
)

func main() {

	r := gin.Default()

	r.GET("/users", userHandler.ShowUserHandler)
	r.POST("/users", userHandler.CreateUserHandler)
	r.GET("/users/:user_id", userHandler.ShowUserByIDHandler)
	r.DELETE("users/:user_id", userHandler.DeleteUserByIDHandler)
	r.PUT("users/:user_id", userHandler.UpdateUserByIDHandler)

	r.Run(":5555")

}
