package handler

import (
	"net/http"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
	"projectpenyewaanlapangan/sportlist"

	"github.com/gin-gonic/gin"
)

type sportlistHandler struct {
	sportlistService sportlist.Service
}

func NewSportListHandler(sportlistService sportlist.Service) *sportlistHandler {
	return &sportlistHandler{sportlistService}
}

//showUserHandler for handling show all user in db from route "/users"
func (h *sportlistHandler) ShowUserHandler(c *gin.Context) {
	sportlist, err := h.sportlistService.GetAllSportList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`message`: `error in internal server error`,
		})
		return
	}
	c.JSON(http.StatusOK, sportlist)
}

// CreateUserHandler for handing if user / external create new user from route "/users"
func (h *sportlistHandler) CreateSportlistHandler(c *gin.Context) {
	var inputSportlist entity.SportListInput

	if err := c.ShouldBindJSON(&inputSportlist); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newSportList, err := h.sportlistService.SaveNewSportList(inputSportlist)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new User", 201, "status Created", newSportList)
	c.JSON(201, response)
}

// get user by 1
// 1. get by id sesuai dengan paramater yg dikasih (repository)
// 2. service akan menampikan hasil user by id dengan format yang sudah ditentukan
// 3. handler kita tangkap id dengan c.Param kemudian kita kirim ke service, terus kita tangkap responsenya

func (h *sportlistHandler) GetSportListByIDHandler(c *gin.Context) {
	id := c.Params.ByName("sportlist_id")

	sportlist, err := h.sportlistService.GetSportListByID(id)
	if err != nil {
		responseError := helper.APIResponse("error bad request sport list ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get sport list by ID", 200, "success", sportlist)
	c.JSON(200, response)
}
