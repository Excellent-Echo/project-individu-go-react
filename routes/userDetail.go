package routes

import (
	"project-individu-go-react/handler"
	userdetail "project-individu-go-react/userDetail"

	"github.com/gin-gonic/gin"
)

var (
	userDetailRepository = userdetail.NewRepository(DB)
	userDetailService    = userdetail.NewService(userDetailRepository)
	userDetailHanlder    = handler.NewUserDetailHandler(userDetailService, authService)
)

func UserDetailRoute(r *gin.Engine) {
	r.GET("/users/user_details", handler.Middleware(userService, authService), userDetailHanlder.GetUserDetailByUserIDHandler)
	r.POST("/users/user_details", handler.Middleware(userService, authService), userDetailHanlder.SaveNewUserDetailHandler)
	r.PUT("/users/user_details", handler.Middleware(userService, authService), userDetailHanlder.UpdateUserDetailByIDHandler)
}
