package handler

import (
	"net/http"
	"project-individu-go-react/config"
	"project-individu-go-react/entity"

	"github.com/gin-gonic/gin"
)

var db = config.Connect()

func GetAllUsers(c *gin.Context) {
	var users []entity.User

	if err := db.Preload("Posts").Preload("Comments").Preload("Likes").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"error message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
