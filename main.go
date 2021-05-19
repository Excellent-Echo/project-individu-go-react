package main

import (
	"project-individu-go-react/config"
	"project-individu-go-react/handler"
	"project-individu-go-react/question"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB = config.Connect()
	userRepository              = user.NewRepository(DB)
	userService                 = user.UserNewService(userRepository)
	userHandler                 = handler.NewUserHandler(userService)
	questionRepository          = question.NewRepository(DB)
	questionService             = question.QuestionNewService(questionRepository)
	questionHandler             = handler.NewQuestionHandler(questionService)
)

func main() {
	r := gin.Default()

	r.GET("/users", userHandler.ShowAllUsersHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	// r.POST("/users/login", userHandler.CreateUserHandler)
	r.GET("/users/:user_id", userHandler.ShowUserByIdHandler)
	r.PUT("/users/:user_id", userHandler.UpdateUserByIDHandler)
	r.DELETE("/users/:user_id", userHandler.DeleteByUserIDHandler)

	r.GET("/questions", questionHandler.ShowAllQuestionsHandler)
	r.POST("/questions/ask", questionHandler.CreateQuestionHandler)
	r.GET("/questions/:id", questionHandler.ShowQuestionByIdHandler)

	r.Run(":4444")
}
