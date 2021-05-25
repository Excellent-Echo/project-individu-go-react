package routes

import (
	"project-individu-go-react/category"
	"project-individu-go-react/handler"

	"github.com/gin-gonic/gin"
)

var (
	categoryRepository = category.NewRepository(DB)
	categoryService    = category.NewService(categoryRepository)
	categoryHandler    = handler.NewCategoryHandler(categoryService)
)

func CategoryRoute(r *gin.Engine) {
	r.GET("/categories", categoryHandler.ShowAllCategoryHandler)
	r.POST("/categories", categoryHandler.CreateCategoryHandler)

}
