package handler

import (
	"net/http"
	"projectpenyewaanlapangan/config"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
	"time"

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

func GetUserByID(c *gin.Context) {
	var user entity.User

	id := c.Param("id_user")

	if err := DB.Where("id = ?", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
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
			"message_error": err.Error(),
		})
		return
	}

	var newUser = entity.User{
		FirstName: getUser.FirstName,
		LastName:  getUser.LastName,
		Email:     getUser.Email,
		Password:  getUser.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newUser)

}

func DeleteByUserID(c *gin.Context) {
	id := c.Param("id")

	var user entity.User

	DB.Where("id = ?", id).Find(&user)

	if user.ID == 0 {
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

func UpdateUserByID(c *gin.Context) {
	var user entity.User
	var userInput entity.UserInput

	id := c.Params.ByName("id")

	if err := DB.Where("id = ?", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
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
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetAllSportList(c *gin.Context) {
	var sportlist []entity.SportList

	if err := DB.Find(&sportlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, sportlist)
}

func GetSportListByID(c *gin.Context) {
	var sportlist entity.SportList

	id := c.Param("id_sportlist")

	if err := DB.Where("id = ?", id).Find(&sportlist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	if sportlist.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "sport list id " + id + " not found in database",
		})
		return
	}

	c.JSON(http.StatusOK, sportlist)
}

// get user by 1
// 1. get by id sesuai dengan paramter yg dikasih (repository)
// 2. service akan menampikan hasil user by id dengan format yang sudah ditentukan
// 3. handler kita tangkap id dengan c.Param kemudian kita kirim ke service, terus kita tangkap responsenya

func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var updateUserInput entity.UpdateUserInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	user, err := h.userService.UpdateUserByID(id, updateUserInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user by ID", 200, "success", user)
	c.JSON(200, response)
}

func (h *userHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := h.userService.DeleteUserByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete user", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success delete user by ID", 200, "success", user)
	c.JSON(200, response)
}
