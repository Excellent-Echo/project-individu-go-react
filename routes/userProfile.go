package routes

import (
	"project-individu-go-react/handler"
	"project-individu-go-react/userProfile"

	"github.com/gin-gonic/gin"
)

var (
	userProfileRepository = userProfile.NewRepository(DB)
	userProfileService    = userProfile.NewService(userProfileRepository)
	userProfileHandler    = handler.NewUserProfileHandler(userProfileService)
)

func UserProfileRoute(r *gin.Engine) {
	r.GET("/users/user_profile", handler.Middleware(userService, authService), userProfileHandler.GetUserProfileByUserIDHandler)
	r.POST("/users/user_profile", handler.Middleware(userService, authService), userProfileHandler.SaveNewUserProfileHandler)
	r.PUT("/users/user_profile", handler.Middleware(userService, authService), userProfileHandler.UpdateUserProfileByIDHandler) // otomatis terhubung ke user yang login
}
