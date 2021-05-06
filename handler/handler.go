package handler

import (
	"net/http"
	"projectpenyewaanlapangan/config"
	"projectpenyewaanlapangan/entity"

	"github.com/gin-gonic/gin"
)

var DB = config.Connect()

func GetAllUser(c *gin.Context) {
	var users []entity.User

	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}
