package main

import (
	"os"
	"project-individu-go-react/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.Static("/images", "./images")

	routes.UserRoute(r)
	routes.QuestionRoute(r)
	// routes.TagRoute(r)
	routes.AnswerRoute(r)
	routes.CategoryRoute(r)
	routes.UserProfileRoute(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
