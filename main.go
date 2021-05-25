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
	// routes.TagRoute(r)
	routes.AnswerRoute(r)
	routes.CategoryRoute(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
