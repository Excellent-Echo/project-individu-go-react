package main

import (
	"projectpenyewaanlapangan/config"
	"projectpenyewaanlapangan/handler"
	"projectpenyewaanlapangan/routes"
	"projectpenyewaanlapangan/sportlist"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.Connect()

	// user entity
	// userRepository = user.NewRepository(DB)
	// userService    = user.NewService(userRepository)
	// userHandler    = handler.NewUserHandler(userService)

	//sportlsit entity
	sportlistRepository = sportlist.NewRepository(DB)
	sportlistService    = sportlist.NewService(sportlistRepository)
	sportlistHandler    = handler.NewSportListHandler(sportlistService)
)

func main() {
	r := gin.Default()

	// endpoint user
	routes.UserRoute(r)

	// endpoint sportlist
	routes.SportListRoutes(r)

	r.Run(":4444")
}
