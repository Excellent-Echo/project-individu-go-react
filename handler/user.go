package handler

import (
	"log"
	"net/http"
	"time"

	"project-individu-go-react/config"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) ShowAllUser(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", http.StatusOK, "error", gin.H{"errors": err.Error()})

		c.JSON(http.StatusOK, responseError)
		return
	}

	response := helper.APIResponse("success get all user", http.StatusOK, "status OK", users)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", http.StatusOK, "bad request", gin.H{"errors": splitError})

		c.JSON(http.StatusOK, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", http.StatusOK, "error", gin.H{"errors": err.Error()})

		c.JSON(http.StatusOK, responseError)
		return
	}
	response := helper.APIResponse("success create new User", http.StatusOK, "Status OK", newUser)
	c.JSON(http.StatusOK, response)
}

var DB = config.Connection()

// GET - ALL - USER
// func GetAllUser(c *gin.Context) {
// 	var users []entity.User

// 	if err := DB.Find(&users).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "error in internal server",
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusOK, users)
// }

// CREATE - NEW - USER
// func CreateNewUser(c *gin.Context) {
// 	var getUser entity.UserInput

// 	if err := c.ShouldBindJSON(&getUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":        "error bad request",
// 			"error_message": err.Error(),
// 		})
// 		return
// 	}

// 	var newUser = entity.User{
// 		First_name: getUser.First_name,
// 		Last_name:  getUser.Last_name,
// 		Password:   getUser.Password,
// 		Email:      getUser.Email,
// 		Created_at: time.Now(),
// 		Updated_at: time.Now(),
// 	}

// 	if err := DB.Create(&newUser).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":        "error bad request",
// 			"error_message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, newUser)

// }

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
