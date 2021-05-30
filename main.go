package main

import (
	"project-individu-go-react/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// CORS
	r.Use(cors.Default())

	routes.UserRoute(r)
	routes.CourseRoute(r)
	r.Run()

}
