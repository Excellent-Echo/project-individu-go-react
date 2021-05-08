package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-individu-go-react/config"
	"project-individu-go-react/entity"
	"time"
)

var DB = config.ConnectToDatabase()

func UpdateUserByID(c *gin.Context) {
	var updateUser entity.User
	var updateUserInput entity.UserInput

	id := c.Params.ByName("user_id")

	if err := DB.Where("id = ?", id).Find(&updateUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	if updateUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	updateUser.RoleID = updateUserInput.RoleID
	updateUser.Firstname = updateUserInput.Firstname
	updateUser.Lastname = updateUserInput.Lastname
	updateUser.Email = updateUserInput.Email
	updateUser.Password = updateUserInput.Password
	updateUser.UpdatedAt = time.Now()

	if err := DB.Where("id = ?", id).Save(&updateUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updateUser)
}
