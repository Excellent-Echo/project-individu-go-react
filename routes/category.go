package routes

import (
	"project-individu-go-react/handler"
	"project-individu-go-react/layer/category"

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
	r.GET("/categories/:category_name", categoryHandler.ShowCategoryByNameHandler)

}
