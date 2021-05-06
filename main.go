package main

import (
	"project-individu-go-react/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", handler.GetAllUsers)
	r.POST("/users", handler.CreateNewUser)
	r.GET("/users/:user_id", handler.GetUserByID)
	r.PUT("/users/:user_id", handler.UpdateUserByID)
	r.DELETE("/users/:user_id", handler.DeleteByUserID)

	r.Run(":4444")
}
