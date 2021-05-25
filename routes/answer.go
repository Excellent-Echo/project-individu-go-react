package routes

import (
	"project-individu-go-react/answer"
	"project-individu-go-react/handler"

	"github.com/gin-gonic/gin"
)

var (
	answerRepository = answer.NewRepository(DB)
	answerService    = answer.NewService(answerRepository)
	answerHandler    = handler.NewAnswerHandler(answerService, authService)
)

func AnswerRoute(r *gin.Engine) {
	r.POST("/questions/:id", handler.Middleware(userService, authService), answerHandler.CreateAnswerHandler)
}
