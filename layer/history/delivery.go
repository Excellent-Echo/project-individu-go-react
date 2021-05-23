package history

import (
	"project-individu-go-react/entities"
	"project-individu-go-react/helper"

	"github.com/gin-gonic/gin"
)

type historyDeliver struct {
	historyService Service
}

func NewHistoryDeliver(historyService Service) *historyDeliver {
	return &historyDeliver{historyService}
}

func (d *historyDeliver) ShowHistoriesDeliver(c *gin.Context) {
	histories, err := d.historyService.GetAllHistory()

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responError)

		return
	}

	response := helper.APIResponse("success get all account", 200, "status OK", histories)
	c.JSON(200, response)
}

func (d *historyDeliver) GetHistoryByHIDDeliver(c *gin.Context) {
	hid := c.Params.ByName("hid")

	history, err := d.historyService.GetHistoryByHID(hid)

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responError)

		return
	}

	response := helper.APIResponse("success get history by hid", 200, "status OK", history)
	c.JSON(200, response)
}

func (d *historyDeliver) GetHistoriesBySIDDeliver(c *gin.Context) {
	sid := c.Params.ByName("sid")

	histories, err := d.historyService.GetHistoriesBySID(sid)

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responError)

		return
	}

	response := helper.APIResponse("success get histories from account sid", 200, "status OK", histories)
	c.JSON(200, response)
}

func (d *historyDeliver) CreateHistoryDeliver(c *gin.Context) {
	var inputHistory entities.InputHistory

	if err := c.ShouldBindJSON(&inputHistory); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("Input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newHistory, err := d.historyService.SaveNewHistory(inputHistory)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
	}

	response := helper.APIResponse("success create new account", 201, "status Created", newHistory)

	c.JSON(201, response)
}
