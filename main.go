package main

import (
	"project-go-react/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", handler.GetAllUser)
	r.POST("/users", handler.CreateNewUser)

	r.Run(":1212")

}
