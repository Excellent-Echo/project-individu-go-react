package main

import (
	"fmt"
	"project-go-react/config"
	"project-go-react/handler"
	"project-go-react/user"

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
	fmt.Println("success connect database")

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userHandler.ShowAllUser)
		v1.GET("/users/:user_id", handler.HandleUsersID)
		v1.POST("/users", handler.HandlePostUser)
		v1.PUT("/users/:user_id", handler.HandlePutUser)
		v1.DELETE("/users/:user_id", handler.HandleDeleteUser)
	}
	router.Run(":8000")

}
