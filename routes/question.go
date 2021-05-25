package routes

import (
	"project-individu-go-react/handler"
	"project-individu-go-react/question"

	"github.com/gin-gonic/gin"
)

var (
	questionRepository = question.NewRepository(DB)
	questionService    = question.QuestionNewService(questionRepository)
	questionHandler    = handler.NewQuestionHandler(questionService, authService)
)

func QuestionRoute(r *gin.Engine) {
	r.GET("/questions", questionHandler.ShowAllQuestionsHandler)
	r.POST("/questions/ask", handler.Middleware(userService, authService), questionHandler.CreateQuestionHandler)
	r.GET("/questions/:id", handler.Middleware(userService, authService), questionHandler.ShowQuestionByIdHandler)
	r.PATCH("/questions/:id/edit", handler.Middleware(userService, authService), questionHandler.UpdateQuestionByIdHandler)

}
