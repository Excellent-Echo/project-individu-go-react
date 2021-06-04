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

	// endpoint booking list

	routes.BookingLIstRoute(r)

	routes.UserProfileRoute(r)

	routes.UserDetailRoute(r)

	r.Run(":4444")
}
