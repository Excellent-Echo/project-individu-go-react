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
	// r.GET("/users/:user_id", handler.GetUserByID)
	// r.POST("/users", handler.CreateNewUser)
	// r.PUT("/users/:user_id", handler.UpdateUserByID)
	// r.DELETE("/users/:user_id", handler.DeleteByUserID)

	// // endpoint sportlist
	// r.GET("/sportlist", handler.GetAllSportList)
	// r.GET("/users/:sportlist_id", handler.GetSportListByID)

	r.Run(":4444")
}
