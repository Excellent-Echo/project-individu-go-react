package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/panduwil/project-individu-go-react/config"
	"github.com/panduwil/project-individu-go-react/entity"

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
		First_name: getUser.First_name,
		Last_name:  getUser.Last_name,
		Password:   getUser.Password,
		Email:      getUser.Email,
		Created_at: time.Now(),
		Updated_at: time.Now(),
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

func HandleUsersID(c *gin.Context) {
	var users entity.User

	id := c.Params.ByName("user_id")

	if err := DB.Where("user_id = ?", id).Find(&users).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, users)
}

func HandleDeleteUser(c *gin.Context) {
	id := c.Params.ByName("user_id")

	if err := DB.Where("user_id = ?", id).Delete(&entity.User{}).Error; err != nil {
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
		"message": "success delete",
	})
}

func HandleUpdateUser(c *gin.Context) {
	var user entity.User
	var userInput entity.UserInput

	id := c.Params.ByName("user_id")
	DB.Where("user_id = ?", id).Find(&user)

	if err := c.ShouldBindJSON(&userInput); err != nil {
		log.Println(err.Error())
		return
	}
	user.First_name = userInput.First_name
	user.Last_name = userInput.Last_name
	user.Email = userInput.Email
	user.Password = userInput.Password
	user.Updated_at = time.Now()

	if err := DB.Where("user_id = ?", id).Save(&user).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
