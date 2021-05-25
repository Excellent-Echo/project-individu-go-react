package main

import (
	"project-individu-go-react/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoute(r)
	routes.QuestionRoute(r)
	routes.TagRoute(r)

	r.Run(":4444")
}
