package routes

import (
	"gorm.io/gorm"
	"project-individu-go-react/auth"
	"project-individu-go-react/config"
	"project-individu-go-react/handler"
	"project-individu-go-react/layer/user"

	"github.com/gin-gonic/gin"
)

var (
	DB             *gorm.DB = config.ConnectToDatabase()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHander(userService, authService)
)

func UserRoute(route *gin.Engine) {
	route.GET("/users", handler.Middleware(userService, authService), userHandler.ShowUserHandler)
	route.POST("/users/register", userHandler.CreateUserHandler)
	route.GET("users/:user_id", handler.Middleware(userService, authService), userHandler.GetUserByIDHandler)
	route.PUT("users/:user_id", handler.Middleware(userService, authService), userHandler.GetandUpdateUserByIDHandler)
	route.DELETE("users/:user_id", handler.Middleware(userService, authService), userHandler.GetandDeleteUserByIDHandler)
	route.POST("/users/login", userHandler.UserLoginHandler)
}
