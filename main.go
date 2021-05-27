package main

import (
	"PROJECT-INDIVIDU-GO-REACT/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", handler.GetAllUser)
	r.GET("/users/:user_id", handler.GetuserByID)
	r.POST("/users", handler.CreateNewUser)
	r.DELETE("/users/:user_id", handler.DeleteByUserID)
	r.POST("/users/:user_id", handler.UpdateUserByID)

	r.Run(":1212")

}
