package main

import (
	"github.com/gin-gonic/gin"
	"project-go-react/handler"
)

func main() {
	r := gin.Default()

	r.GET("/users", handler.GetAllUser)
	r.POST("/users", handler.CreateNewUser)

	r.Run(":1212")

}
