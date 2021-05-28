package main

import (
	"project-individu-go-react/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	route.UserRoutes(r)

	r.Run()
}
