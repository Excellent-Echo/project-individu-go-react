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
