package main

import (
	"project-individu-go-react/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("users", handler.GetAllUsers)

	r.Run(":4444")
}
