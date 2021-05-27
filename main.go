package main

import (
	"PROJECT-INDIVIDU_GO_REACT/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", handler.GetAllUser)
	r.GET("/users/:user_id", handler.HandleUsersID)
	r.POST("/users", handler.CreateNewUser)
	r.DELETE("/users/:user_id", handler.HandleDeleteUser)
	r.POST("/users/:user_id", handler.HandleUpdateUser)

	r.Run(":1212")

}
