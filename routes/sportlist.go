package routes

import (
	"projectpenyewaanlapangan/handler"
	"projectpenyewaanlapangan/sportlist"

	"github.com/gin-gonic/gin"
)

var (
	// DB          *gorm.DB = config.Connect()
	sportlistRepository = sportlist.NewRepository(DB)
	sportlistService    = sportlist.NewService(sportlistRepository)
	sportlistHandler    = handler.NewSportListHandler(sportlistService)
)

func SportListRoutes(r *gin.Engine) {
	r.GET("/sportlist", sportlistHandler.ShowUserHandler)
	r.POST("/sportlist/register", sportlistHandler.CreateSportlistHandler)
	r.GET("/sportlist/:sportlist_id", sportlistHandler.GetSportListByIDHandler)
}
