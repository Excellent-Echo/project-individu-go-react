package handler

import (
	"net/http"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/fieldlist"
	"projectpenyewaanlapangan/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type fieldlistHandler struct {
	fieldlistService fieldlist.Service
}

func NewFieldListHandler(fieldlistService fieldlist.Service) *fieldlistHandler {
	return &fieldlistHandler{fieldlistService}
}

//showFieldListHandler for handling show all fieldlist in db from route "/fieldlist"
func (h *fieldlistHandler) showFieldListHandler(c *gin.Context) {
	fieldlist, err := h.fieldlistService.GetAllFieldList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`message`: `error in internal server error`,
		})
	}
	c.JSON(http.StatusOK, fieldlist)
}

// CreateFieldListHandler for handing if field / external create new fieldlist from route "/fieldlist"
func (h *fieldlistHandler) CreateFieldListHandler(c *gin.Context) {
	sportlistData := int(c.MustGet("currentSportList").(int))

	if sportlistData == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}

	sportlistID := strconv.Itoa(sportlistData)

	var inputFieldList entity.FieldListInput

	if err := c.ShouldBindJSON(&inputFieldList); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	NewFieldList, err := h.fieldlistService.SaveNewFieldList(inputFieldList, sportlistID)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new Field List", 201, "status Created", NewFieldList)
	c.JSON(201, response)
}

func (h *fieldlistHandler) GetFieldListByID(c *gin.Context) {
	id := c.Params.ByName("fieldlistid")

	fieldlist, err := h.fieldlistService.GetFieldListByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request field list ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get field list by ID", 200, "success", fieldlist)
	c.JSON(200, response)
}
