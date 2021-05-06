package main

import (
	"projectpenyewaanlapangan/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// endpoint user
	r.GET("/users", handler.GetAllUser)
	r.GET("/users/:user_id", handler.GetUserByID)
	r.POST("/users", handler.CreateNewUser)
	r.PUT("/users/:user_id", handler.UpdateUserByID)
	r.DELETE("/users/:user_id", handler.DeleteByUserID)

	// endpoint sportlist
	r.GET("/sportlist", handler.GetAllSportList)
	r.GET("/users/:sportlist_id", handler.GetSportListByID)

	r.Run(":4444")
}
