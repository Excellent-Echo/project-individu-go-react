package routes

import (
	"project-individu-go-react/handler"
	"project-individu-go-react/layer/tag"

	"github.com/gin-gonic/gin"
)

var (
	tagRepository = tag.NewRepository(DB)
	tagService    = tag.NewService(tagRepository)
	tagHandler    = handler.NewTagHandler(tagService)
)

func TagRoute(r *gin.Engine) {
	r.GET("/tags", tagHandler.ShowAllTagsHandler)
	r.POST("/tags", tagHandler.CreateTagHandler)

}
