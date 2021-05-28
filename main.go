package main

import (
	"projectpenyewaanlapangan/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// endpoint user
	routes.UserRoute(r)

	// endpoint sportlist
	routes.SportListRoutes(r)

	r.Run()
}
