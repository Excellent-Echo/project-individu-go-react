package main

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/routes"
)

func main() {
	router := gin.Default()

	routes.UserRoute(router)
	//routes.PsikologRoute(router)
	//routes.RoleRoute(router)
	//routes.BookingDetailRoute(router)
	//routes.BookingRoute(router)

	router.Run()

}
