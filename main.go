package main

import (
	"projectpenyewaanlapangan/config"
	"projectpenyewaanlapangan/handler"
	"projectpenyewaanlapangan/user"

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

	// endpoint user
	r.GET("/users", userHandler.ShowUserHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.GET("/users/:user_id", userHandler.GetUserByIDHandler)
	r.DELETE("/users/:user_id", userHandler.DeleteUserByIDHandler)
	r.PUT("/users/:user_id", userHandler.UpdateUserByIDHandler)

	// // endpoint sportlist
	// r.GET("/sportlist", handler.GetAllSportList)
	// r.GET("/users/:sportlist_id", handler.GetSportListByID)

	r.Run(":4444")
}
