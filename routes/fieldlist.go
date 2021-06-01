package routes

import (
	"projectpenyewaanlapangan/fieldlist"
	"projectpenyewaanlapangan/handler"

	"github.com/gin-gonic/gin"
)

var (
	fieldlistRepository = fieldlist.NewRepository(DB)
	fieldlistService    = fieldlist.NewService(fieldlistRepository)
	fieldlistHandler    = handler.NewFieldListHandler(fieldlistService)
)

func FieldListRoutes(r *gin.Engine) {
	r.GET("/fieldlist", fieldlistHandler.ShowFieldListHandler)
	r.POST("/fieldlist", fieldlistHandler.SaveNewFieldListHandler)
	r.GET("/fieldlist/:field_id", fieldlistHandler.GetFieldListByID)
	r.PUT("/fieldlist/:field_id", fieldlistHandler.UpdateFieldListByIdHandler)
}
