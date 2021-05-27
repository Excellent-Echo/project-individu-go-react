package handler

import (
	"net/http"
	"project-go-react/config"
	"project-go-react/entity"
	"time"

	"github.com/gin-gonic/gin"
)

var DB = config.Connection()

func GetAllUser(c *gin.Context) {
	var users []entity.User

	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in internal server",
		})

		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	var user entity.User

	id := c.Param("user_id")

	if err := DB.Where("id = ?", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"error_message": err.Error(),
		})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"error_message": "user id" + id + "not found in database",
		})
		return
	}

	c.JSON(http.StatusOK, user)
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
		Password:  getUser.Password,
		Email:     getUser.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newUser)

}

func DeleteByUserID(c *gin.Context) {
	id := c.Param("user_id")

	var user entity.User

	DB.Where("id = ?", id).Find(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"error_message": "user id" + id + "not found in database",
		})
		return
	}

	if err := DB.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error bad request",
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        "delete success",
		"error_message": "user id" + id + "successfull delete",
	})
}

func UpdateByUserID(c *gin.Context) {
	var user entity.User
	var userInput entity.UserInput

	id := c.Params.ByName("user_id")

	if err := DB.Where("id = ?", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error bad request",
			"error_message": err.Error(),
		})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"error_message": "user id" + id + "not found in database",
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
	user.Email = userInput.Email
	user.Password = userInput.Password
	user.UpdatedAt = time.Now()

	if err := DB.Where("id = ?", id).Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error bad request",
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)

}
