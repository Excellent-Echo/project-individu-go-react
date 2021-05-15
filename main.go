package main

import (
	"projectpenyewaanlapangan/config"
	"projectpenyewaanlapangan/handler"
	"projectpenyewaanlapangan/sportlist"
	"projectpenyewaanlapangan/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.Connect()

	// user entity
	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)
	userHandler    = handler.NewUserHandler(userService)

	//sportlsit entity
	sportlistRepository = sportlist.NewRepository(DB)
	sportlistService    = sportlist.NewService(sportlistRepository)
	sportlistHandler    = handler.NewSportListHandler(sportlistService)
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
	r.GET("/sportlist", sportlistHandler.ShowUserHandler)
	r.POST("/sportlist/register", sportlistHandler.CreateSportlistHandler)
	r.GET("/users/:sportlist_id", sportlistHandler.GetSportListByIDHandler)

	r.Run(":4444")
}
