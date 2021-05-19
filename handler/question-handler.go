package handler

import (
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/question"

	"github.com/gin-gonic/gin"
)

type questionHandler struct {
	questionService question.QuestionService
}

func NewQuestionHandler(questionService question.QuestionService) *questionHandler {
	return &questionHandler{questionService}
}

func (h *questionHandler) ShowAllQuestions(c *gin.Context) {
	questions, err := h.questionService.FindAllQuestions()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	userResponse := helper.APIResponse("get all questions succeed", 200, "success", questions)
	c.JSON(200, userResponse)

}

func (h *questionHandler) CreateQuestionHandler(c *gin.Context) {
	var inputQuestion entity.QuestionInput

	if err := c.ShouldBindJSON(&inputQuestion); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	response, err := h.questionService.SaveNewQuestion(inputQuestion)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	userResponse := helper.APIResponse("insert new question succeed", 201, "success", response)
	c.JSON(201, userResponse)
}
