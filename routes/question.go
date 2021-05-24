package routes

import (
	"project-individu-go-react/handler"
	"project-individu-go-react/question"

	"github.com/gin-gonic/gin"
)

var (
	questionRepository = question.NewRepository(DB)
	questionService    = question.QuestionNewService(questionRepository)
	questionHandler    = handler.NewQuestionHandler(questionService)
)

func QuestionRoute(r *gin.Engine) {
	r.GET("/questions", questionHandler.ShowAllQuestionsHandler)
	r.POST("/questions/ask", questionHandler.CreateQuestionHandler)
	r.GET("/questions/:id", questionHandler.ShowQuestionByIdHandler)

}
