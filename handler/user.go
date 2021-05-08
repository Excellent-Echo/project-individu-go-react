package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-individu-go-react/config"
	"project-individu-go-react/entity"
	"time"
)

var DB = config.ConnectToDatabase()

func GetUserByID(c *gin.Context) {
	var GetUserID entity.User

	id := c.Param("user_id")

	if err := DB.Where("id = ?", id).Preload("Booking").Find(&GetUserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	if GetUserID.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	c.JSON(http.StatusOK, GetUserID)
}

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

func DeleteByUserID(c *gin.Context) {
	id := c.Param("user_id")

	var DeleteUser entity.User

	DB.Where("id = ?", id).Find(&DeleteUser)

	if DeleteUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := DB.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "delete success",
		"message": "user id " + id + " successfull delete",
	})
}
