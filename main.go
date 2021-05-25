package main

import (
	"project-individu-go-react/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	route.UserRoutes(r)

	r.Run()
}
