package routes

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/handler"
	"project-individu-go-react/layer/role"
)

var (
	roleRepository = role.NewRepository(DB)
	roleService    = role.NewService(roleRepository)
	roleHandler    = handler.NewRoleHander(roleService, authService)
)

func RoleRoute(router *gin.Engine) {
	router.GET("/roles", handler.Middleware(userService, authService), roleHandler.ShowRoleHandler)
	router.POST("/roles", roleHandler.CreateRoleHandler)
	router.GET("roles/:role_id", handler.Middleware(userService, authService), roleHandler.GetRoleByIDHandler)
	router.PUT("roles/:role_id", handler.Middleware(userService, authService), roleHandler.GetandUpdateRoleByIDHandler)
	router.DELETE("roles/:role_id", handler.Middleware(userService, authService), roleHandler.GetandDeleteRoleByIDHandler)
}
