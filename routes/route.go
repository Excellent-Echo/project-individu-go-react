package routes

import (
	"project-individu-go-react/auth"
	"project-individu-go-react/config"
	course "project-individu-go-react/courses"
	"project-individu-go-react/handler"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
)

var (
	DB               = config.Connection()
	userRepository   = user.NewRepository(DB)
	userService      = user.NewService(userRepository)
	authService      = auth.NewService()
	userHandler      = handler.NewUserHandler(userService, authService)
	courseRepository = course.NewRepository(DB)
	courseService    = course.NewService(courseRepository)
	courseHandler    = handler.NewCourseHandler(courseService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", handler.Middleware(userService, authService), userHandler.ShowUserHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.GET("/users/:user_id", handler.Middleware(userService, authService), userHandler.ShowUserByIDHandler)
	r.DELETE("users/:user_id", handler.Middleware(userService, authService), userHandler.DeleteUserByIDHandler)
	r.PUT("users/:user_id", handler.Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
}

func CourseRoute(r *gin.Engine) {
	r.GET("/courses", courseHandler.ShowCourseHandler)
}
