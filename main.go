package main

import (
	"os"
	"project-individu-go-react/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoute(r)
	routes.QuestionRoute(r)
	routes.TagRoute(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
