package main

import (
	"project-individu-go-react/auth"
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
	authService                 = auth.NewService()
	userHandler                 = handler.NewUserHandler(userService, authService)
	questionRepository          = question.NewRepository(DB)
	questionService             = question.QuestionNewService(questionRepository)
	questionHandler             = handler.NewQuestionHandler(questionService)
)

func main() {
	r := gin.Default()

	r.POST("/users/login", userHandler.LoginUserHandler)
	r.GET("/users", handler.Middleware(userService, authService), userHandler.ShowAllUsersHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	// r.POST("/users/login", userHandler.CreateUserHandler)
	r.GET("/users/:id", handler.Middleware(userService, authService), userHandler.ShowUserByIdHandler)
	r.PUT("/users/:id", handler.Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
	r.DELETE("/users/:id", handler.Middleware(userService, authService), userHandler.DeleteByUserIDHandler)

	r.GET("/questions", questionHandler.ShowAllQuestionsHandler)
	r.POST("/questions/ask", questionHandler.CreateQuestionHandler)
	r.GET("/questions/:id", questionHandler.ShowQuestionByIdHandler)

	r.Run(":4444")
}
