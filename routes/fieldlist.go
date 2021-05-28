package routes

import (
	"projectpenyewaanlapangan/fieldlist"
	"projectpenyewaanlapangan/handler"

	"github.com/gin-gonic/gin"
)


var (
	fieldlistRepository = fieldlist.NewRepository(DB)
	fieldlistService = fieldlist.NewService(fieldlistRepository)
	fieldlistHandler = handler.NewFieldListHandler(fieldlistService)
)

func FieldListRoutes(r *gin.Engine) {
	r.GET("/fieldlist", fieldlistHandler.showFieldListHandler)
	r.POST("/fieldlist/register", fieldlistHandler.CreateFieldListHandler)
	r.GET("/fieldlist/:fieldlistbyid", fieldlistHandler.)
}