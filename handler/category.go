package handler

import (
	"project-individu-go-react/category"
	"project-individu-go-react/helper"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService category.Service
}

func NewCategoryHandler(categoryService category.Service) *CategoryHandler {
	return &CategoryHandler{categoryService}
}

func (h *CategoryHandler) ShowAllCategoryHandler(c *gin.Context) {
	categories, err := h.categoryService.GetAllCategory()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all categories", 200, "success", categories)
	c.JSON(200, response)
}

func (h *CategoryHandler) ShowCategoryByIDHandler(c *gin.Context) {
	id := c.Param("id")

	category, err := h.categoryService.GetCategoryByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get category by ID", 200, "success", category)
	c.JSON(200, response)
}
