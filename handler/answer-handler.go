package handler

import (
	"project-individu-go-react/answer"
	"project-individu-go-react/auth"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type answerHandler struct {
	answerService answer.AnswerService
	authService   auth.Service
}

func NewAnswerHandler(answerService answer.AnswerService, authService auth.Service) *answerHandler {
	return &answerHandler{answerService, authService}
}

func (h *answerHandler) CreateAnswerHandler(c *gin.Context) {
	var inputAnswer entity.AnswerInput

	if err := c.ShouldBindJSON(&inputAnswer); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userID := int(c.MustGet("currentUser").(int))
	id := c.Params.ByName("id")
	questionID, _ := strconv.Atoi(id)

	response, err := h.answerService.PostNewAnswer(questionID, userID, inputAnswer)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	newResponse := helper.APIResponse("insert new answer succeed", 201, "success", response)
	c.JSON(201, newResponse)
}
