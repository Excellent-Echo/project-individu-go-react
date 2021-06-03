package handler

import (
	"log"
	"net/http"
	"project-go-react/config"
	"project-go-react/entity"

	"github.com/gin-gonic/gin"
)

var db = config.Connection()

func HandleUsers(c *gin.Context) {
	var users []entity.User

	err := db.Preload("Events").Find(&users).Error

	if err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func HandleUsersID(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("user_id")
	if err := db.Where("id = ?", id).Preload("Events").Find(&user).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func HandlePostUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		return
	}
	if err := db.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func HandleDeleteUser(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("user_id")

	if err := db.Delete(&user, id).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "success delete",
	})
}

func HandlePutUser(c *gin.Context) {
	var user entity.User
	var userInput entity.User

	id := c.Params.ByName("user_id")
	db.Where("id = ?", id).Find(&user)

	if err := c.ShouldBindJSON(&userInput); err != nil {
		log.Println(err.Error())
		return
	}

	user.FirstName = userInput.FirstName
	user.LastName = userInput.LastName
	user.Email = userInput.Email

	if err := db.Where("id = ?", id).Save(&user).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
