package routes

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/handler"
	"project-individu-go-react/layer/userprofile"
)

var (
	userProfileRepository = userprofile.NewRepository(DB)
	userProfileService    = userprofile.NewService(userProfileRepository)
	userProfileHandler    = handler.NewUserProfileHandler(userProfileService)
)

func UserProfileRoute(router *gin.Engine) {
	router.GET("/user_profile", handler.Middleware(userService, authService), userProfileHandler.GetUserProfileByUserIDHandler)
	router.POST("/user_profile", handler.Middleware(userService, authService), userProfileHandler.SaveNewUserProfileHandler)
	router.PUT("/user_profile", handler.Middleware(userService, authService), userProfileHandler.UpdateUserProfileByIDHandler)
}
