package routes

import (
	"doctorsAppointment/auth"
	"doctorsAppointment/config"
	"doctorsAppointment/handler"
	"doctorsAppointment/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", handler.Middleware(userService, authService), userHandler.ShowUserHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
}
