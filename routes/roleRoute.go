package routes

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/handler"
)

func RoleRoute(router *gin.Engine) {
	router.GET("/roles", handler.GetAllRole)
	router.POST("/roles", handler.CreateNewRole)
	router.GET("roles/:role_id", handler.GetRoleByID)
	router.PUT("roles/:role_id", handler.UpdateRoleByID)
	router.DELETE("roles/:role_id", handler.DeleteByRoleID)
}
