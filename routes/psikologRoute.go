package routes

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/handler"
	"project-individu-go-react/layer/psikolog"
)

var (
	psikologRepository = psikolog.NewRepository(DB)
	psikologService    = psikolog.NewService(psikologRepository)
	psikologHandler    = handler.NewPsikologHandler(psikologService, authService)
)

func PsikologRoute(router *gin.Engine) {
	router.GET("/psikologs", handler.Middleware(userService, authService), psikologHandler.ShowPsikologHandler)
	router.POST("/psikologs", psikologHandler.CreatePsikologHandler)
	router.GET("psikologs/:psikolog_id", handler.Middleware(userService, authService), psikologHandler.GetPsikologByIDHandler)
	router.PUT("psikologs/:psikolog_id", handler.Middleware(userService, authService), psikologHandler.UpdatePsikologByIDHandler)
	router.DELETE("psikologs/:psikolog_id", handler.Middleware(userService, authService), psikologHandler.DeletePsikologByIDHandler)
}
