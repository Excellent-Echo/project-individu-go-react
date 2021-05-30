package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"project-go-react/auth"
	"project-go-react/config"
	"project-go-react/handler"
	"project-go-react/user"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func main() {
	r := gin.Default()

	r.GET("/users", handler.MiddleWare(userService, authService), userHandler.ShowUserHandler)
	r.GET("/users/:user_id", handler.MiddleWare(userService, authService), userHandler.GetUserByIDHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.PUT("/users/:user_id", handler.MiddleWare(userService, authService), userHandler.UpdateUserByIDHandler)
	r.DELETE("/users/:user_id", handler.MiddleWare(userService, authService), userHandler.DeleteUserByIDHandler)

	r.Run(":1212")

}
