package routes

import (
	"projectpenyewaanlapangan/handler"
	"projectpenyewaanlapangan/userprofile"

	"github.com/gin-gonic/gin"
)

var (
	userProfileRepository = userprofile.NewRepository(DB)
	userProfileService    = userprofile.NewService(userProfileRepository)
	userProfileHandler    = handler.NewUserProfileHandler(userProfileService)
)

func UserProfileRoute(r *gin.Engine) {
	r.GET("/user_profile", handler.Middleware(userService, authService), userProfileHandler.GetUserProfileByUserIDHandler)
	r.POST("/user_profile", handler.Middleware(userService, authService), userProfileHandler.SaveNewUserProfileHandler)
	r.PUT("/user_profile", handler.Middleware(userService, authService), userProfileHandler.UpdateUserProfileByIDHandler) // otomatis terhubung ke user yang login
}
