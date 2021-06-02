package handler

import (
	"fmt"
	"net/http"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/fieldlist"
	"projectpenyewaanlapangan/helper"

	"github.com/gin-gonic/gin"
)

type fieldlistHandler struct {
	fieldlistService fieldlist.Service
}

func NewFieldListHandler(fieldlistService fieldlist.Service) *fieldlistHandler {
	return &fieldlistHandler{fieldlistService}
}

// showFieldListHandler for handling show all fieldlist in db from route "/fieldlist"
func (h *fieldlistHandler) ShowFieldListHandler(c *gin.Context) {
	fieldlist, err := h.fieldlistService.GetAllFieldList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`message`: `error in internal server error`,
		})
	}
	c.JSON(http.StatusOK, fieldlist)
}

func (h *fieldlistHandler) GetFieldListByID(c *gin.Context) {
	id := c.Params.ByName("field_id")

	fieldlist, err := h.fieldlistService.GetFieldListByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request field list ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get field list by ID", 200, "success", fieldlist)
	c.JSON(200, response)
}

func (h *fieldlistHandler) SaveNewFieldListHandler(c *gin.Context) {

	var fieldListInput entity.FieldListInput
	file, err := c.FormFile("field_image") // postman

	if err != nil {
		responseError := helper.APIResponse("status bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	path := fmt.Sprintf("images/field_image-%v-%s", fieldListInput, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		// log.Println("error line 63")
		responseError := helper.APIResponse("status bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	pathFieldSave := "https://rentfield.herokuapp.com/" + path

	newField, err := h.fieldlistService.SaveNewFieldList(pathFieldSave, fieldListInput)

	if err != nil {
		responseError := helper.APIResponse("Internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create user profile", 201, "success", newField)
	c.JSON(201, response)
}

func (h *fieldlistHandler) UpdateFieldListByIdHandler(c *gin.Context) {
	id := c.Params.ByName("field_id")

	var updateFieldInput entity.FieldListInput

	file, err := c.FormFile("field_image")

	if err != nil {
		responseError := helper.APIResponse("status bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	path := fmt.Sprintf("images/field_image-%v-%s", updateFieldInput, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		// log.Println("error line 63")
		responseError := helper.APIResponse("status bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	pathFieldSave := "https://rentfield.herokuapp.com/" + path

	updateField, err := h.fieldlistService.UpdateFieldListById(pathFieldSave, id, updateFieldInput)

	if err != nil {
		responseError := helper.APIResponse("Internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user profile", 200, "success", updateField)
	c.JSON(200, response)

}
