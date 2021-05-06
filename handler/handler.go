package handler

import (
	"net/http"
	"project-individu-go-react/config"
	"project-individu-go-react/entity"
	"time"

	"github.com/gin-gonic/gin"
)

var db = config.Connect()

func GetAllUsers(c *gin.Context) {
	var users []entity.User

	if err := db.Preload("Posts").Preload("Comments").Preload("Likes").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateNewUser(c *gin.Context) {
	var getUser entity.UserInput

	if err := c.ShouldBindJSON(&getUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"error_message": err.Error(),
		})
		return
	}

	var newUser = entity.User{
		FirstName: getUser.FirstName,
		LastName:  getUser.LastName,
		UserName:  getUser.UserName,
		Email:     getUser.Email,
		Password:  getUser.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newUser)

}

func GetUserByID(c *gin.Context) {
	var user entity.User

	id := c.Param("user_id")

	if err := db.Where("user_id = ?", id).Preload("Posts").Preload("Comments").Preload("Likes").Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"error_message": err.Error(),
		})
		return
	}

	if user.UserID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "data not found",
			"error_message": "user id " + id + " not found in database",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserByID(c *gin.Context) {
	var user entity.User
	var userInput entity.UserInput

	id := c.Params.ByName("user_id")

	if err := db.Where("user_id = ?", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"error_message": err.Error(),
		})
		return
	}

	if user.UserID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "data not found",
			"error_message": "user id " + id + " not found in database",
		})
		return
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"error_message": err.Error(),
		})
		return
	}

	user.FirstName = userInput.FirstName
	user.LastName = userInput.LastName
	user.UserName = userInput.UserName
	user.Email = userInput.Email
	user.Password = userInput.Password
	user.UpdatedAt = time.Now()

	if err := db.Where("user_id = ?", id).Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteByUserID(c *gin.Context) {
	id := c.Param("user_id")

	var user entity.User

	db.Where("user_id = ?", id).Find(&user)

	if user.UserID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"error_message": "user id " + id + " not found in database",
		})
		return
	}

	if err := db.Where("user_id = ?", id).Delete(&entity.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "deleted succeed",
		"message": "user id " + id + " deleted successfully",
	})
}
