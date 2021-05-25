package handler

import (
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/tag"

	"github.com/gin-gonic/gin"
)

type tagHandler struct {
	tagService tag.TagService
}

func NewTagHandler(tagService tag.TagService) *tagHandler {
	return &tagHandler{tagService}
}

func (h *tagHandler) ShowAllTagsHandler(c *gin.Context) {
	tags, err := h.tagService.GetAllTags()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("get all tags succeed", 200, "success", tags)
	c.JSON(200, response)

}

func (h *tagHandler) CreateTagHandler(c *gin.Context) {
	var tagInput entity.TagInput

	if err := c.ShouldBindJSON(&tagInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	response, err := h.tagService.SaveNewTag(tagInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	newResponse := helper.APIResponse("insert new tag succeed", 201, "success", response)
	c.JSON(201, newResponse)
}
