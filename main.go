package main

import (
	"project-individu-go-react/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/kontak", handler.GetAllKontak)
	r.GET("/kontak/:idKontak", handler.HandleIDKontak)
	r.POST("/kontak", handler.CreateNewKontak)
	r.DELETE("/kontak/:idKontak", handler.HandleDelKontak)
	r.GET("/kontak/:idKontak", handler.HandleUpKontak)

	r.Run(":8080")
}
