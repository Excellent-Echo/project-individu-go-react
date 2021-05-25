package main

import (
	"project-individu-go-react/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.UserRoute(r)

	r.Run()

}
