package routes

import (
	"project-individu-go-react/category"
	"project-individu-go-react/handler"

	"github.com/gin-gonic/gin"
)

var (
	categoryRepository = category.NewRepository(DB)
	categoryService    = category.NewService(categoryRepository, courseRepository)
	categoryHandler    = handler.NewCategoryHandler(categoryService)
)

func CategoryRoute(r *gin.Engine) {
	r.GET("/category", categoryHandler.ShowAllCategoryHandler)
	r.GET("/category/:id", categoryHandler.ShowCategoryByIDHandler)
}
