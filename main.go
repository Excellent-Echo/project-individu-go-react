package main

import (
	"projectpenyewaanlapangan/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/images", "./images")

	// endpoint user
	routes.UserRoute(r)

	// endpoint fieldlist
	routes.FieldListRoutes(r)

	r.Run(":4444")
}
