package routes

import (
	"project-individu-go-react/handler"
	"project-individu-go-react/layer/answer"

	"github.com/gin-gonic/gin"
)

var (
	answerRepository = answer.NewRepository(DB)
	answerService    = answer.NewService(answerRepository)
	answerHandler    = handler.NewAnswerHandler(answerService, authService)
)

func AnswerRoute(r *gin.Engine) {
	r.POST("/questions/:id", handler.Middleware(userService, authService), answerHandler.CreateAnswerHandler)
	r.PATCH("/answers/:id", handler.Middleware(userService, authService), answerHandler.UpdateAnswerHandler)
	r.DELETE("/answers/:id", handler.Middleware(userService, authService), answerHandler.DeleteAnswerHandler)
}
