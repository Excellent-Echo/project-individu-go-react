package routes

import (
	"project-individu-go-react/auth"
	"project-individu-go-react/config"
	"project-individu-go-react/handler"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connect()
	userRepository          = user.NewRepository(DB)
	userService             = user.UserNewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", handler.Middleware(userService, authService), userHandler.ShowAllUsersHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.GET("/users/:id", handler.Middleware(userService, authService), userHandler.ShowUserByIdHandler)
	r.PUT("/users/:id", handler.Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
	r.DELETE("/users/:id", handler.Middleware(userService, authService), userHandler.DeleteByUserIDHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
}
