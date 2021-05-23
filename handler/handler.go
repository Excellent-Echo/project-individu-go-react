// package handler

// import (
// 	"PROJECT-INDIVIDU-GO-REAC/entity"
// 	"PROJECT-INDIVIDU-GO-REACT/config"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// var DB = config.Connect()

// func GetuserByID(c *gin.Context) {
// 	var user entity.User

// 	id := c.Params("user_id")

// 	if err := DB.Where("id = ?", id).Preload("UserDetail").Preload("Mentor").Find(&user).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":        "error bad request",
// 			"message_error": err.Error(),
// 		})
// 		return
// 	}

// 	if user.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":        "error not found",
// 			"message_error": "user id " + id + " not found in database",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

// func DeleteByUserID(c *gin.Context) {
// 	id := c.Param("user_id")

// 	var user entity.User

// 	DB.Where("id = ?", id).Find(&user)

// 	if user.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":        "error not found",
// 			"message_error": "user id " + id + " not found in database",
// 		})
// 		return
// 	}

// 	if err := DB.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":        "error in internal server",
// 			"message_error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  "delete success",
// 		"message": "user id " + id + " successfull delete",
// 	})
// }

// func UpdateUserByID(c *gin.Context) {
// 	var user entity.User
// 	var userInput entity.UserInput

// 	id := c.Params.ByName("user_id")

// 	if err := DB.Where("id = ?", id).Find(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":        "error in internal server",
// 			"message_error": err.Error(),
// 		})
// 		return
// 	}

// 	if user.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":        "error not found",
// 			"message_error": "user id " + id + " not found in database",
// 		})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&userInput); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":        "error bad request",
// 			"message_error": err.Error(),
// 		})
// 		return
// 	}

// 	user.FirstName = userInput.FirstName
// 	user.LastName = userInput.LastName
// 	user.Email = userInput.Email
// 	user.Password = userInput.Password
// 	user.UpdatedAt = time.Now()

// 	if err := DB.Where("id = ?", id).Save(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":        "error in internal server",
// 			"message_error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }
// }
