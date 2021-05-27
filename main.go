package main

import (
	"github.com/gin-gonic/gin"
	"project-go-react/handler"
)

func main() {
	r := gin.Default()

	r.GET("/users", handler.GetAllUser)
	r.GET("/users/:user_id", handler.GetUserByID)
	r.POST("/users", handler.CreateNewUser)
	r.PUT("/users/:user_id", handler.UpdateByUserID)
	r.DELETE("/users/:user_id", handler.DeleteByUserID)

	r.Run(":1212")

}
