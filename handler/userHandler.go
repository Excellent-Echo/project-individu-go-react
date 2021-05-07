package handler

import (
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

func (h *userHandler) ShowAllUsersHandler(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	userResponse := helper.APIResponse("get all users succeed", 200, "success", users)
	c.JSON(200, userResponse)
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	response, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	userResponse := helper.APIResponse("insert user data succeed", 201, "success", response)
	c.JSON(201, userResponse)
}

func (h *userHandler) ShowUserByIdHandler(c *gin.Context) {
	id := c.Param("user_id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userResponse := helper.APIResponse("get user succeed", 200, "success", user)
	c.JSON(200, userResponse)
}

// func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
// 	var userInput entity.UserInput

// 	id := c.Params.ByName("user_id")
// 	idNumber, err := strconv.Atoi(id)

// 	if err != nil || idNumber == 0 {
// 		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

// 		c.JSON(400, responseError)
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&userInput); err != nil {
// 		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

// 		c.JSON(400, responseError)
// 		return
// 	}

// 	user, err := h.userService.UpdateUserByID(idNumber)

// 	if err != nil {
// 		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

// 		c.JSON(500, responseError)
// 		return
// 	}

// 	userResponse := helper.APIResponse("update user succeed", 200, "success", user)
// 	c.JSON(200, userResponse)

// }

// func (h *userHandler) DeleteByUserIDHandler(c *gin.Context) {
// 	id := c.Param("user_id")
// 	idNumber, err := strconv.Atoi(id)

// 	if err != nil || idNumber == 0 {
// 		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

// 		c.JSON(400, responseError)
// 		return
// 	}

// 	user, err := h.userService.DeleteByUserID(idNumber)

// 	if err != nil {
// 		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

// 		c.JSON(500, responseError)
// 		return
// 	}

// 	userResponse := helper.APIResponse("user was deleted successfully", 200, "success", user)
// 	c.JSON(200, userResponse)

// }
