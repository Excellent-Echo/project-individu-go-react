package handler

import (
	"project-individu-go-react/auth"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// function handler for Get All data User
func (h *userHandler) ShowUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		responseErr := helper.APIResponse("internal server error", 500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}

	response := helper.APIResponse("success get all data user", 200, "status ok", users)

	c.JSON(200, response)
}

//function handler for create new user
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitErr := helper.SplitErrorInformation(err)
		responseErr := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitErr})
		c.JSON(400, responseErr)
		return
	}

	newUser, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		responseErr := helper.APIResponse("internal server error", 500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}
	response := helper.APIResponse("success create new user", 201, "status created", newUser)
	c.JSON(201, response)
}

//function handler for get data user by id
func (h *userHandler) ShowUserByIDHandler(c *gin.Context) {
	id := c.Param("user_id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		responseErr := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseErr)
		return
	}
	response := helper.APIResponse("success get data user by id", 200, "success", user)
	c.JSON(200, response)
}

//function handler for delete user by id
func (h *userHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := c.Param("user_id")

	user, err := h.userService.DeleteByUserID(id)

	if err != nil {
		responseErr := helper.APIResponse("error bad request delete user", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseErr)
		return
	}
	response := helper.APIResponse("success delete user", 200, "success", user)
	c.JSON(200, response)
}

func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Param("user_id")

	var updateUserInput entity.UpdateUserInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		splitErr := helper.SplitErrorInformation(err)
		responseErr := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": splitErr})

		c.JSON(400, responseErr)
		return
	}

	user, err := h.userService.UpdateUserByID(id, updateUserInput)

	if err != nil {
		responseErr := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseErr)
		return
	}

	response := helper.APIResponse("success update user by id", 200, "success", user)
	c.JSON(200, response)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser entity.LoginUserInput

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		splitErr := helper.SplitErrorInformation(err)
		responseErr := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": splitErr})

		c.JSON(400, responseErr)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseErr := helper.APIResponse("input data error", 401, "bad request", gin.H{"error": err})

		c.JSON(400, responseErr)
		return
	}

	token, err := h.authService.GenerateToken(userData.ID)
	if err != nil {
		responseErr := helper.APIResponse("input data error", 401, "bad request", gin.H{"error": err})

		c.JSON(400, responseErr)
		return
	}

	response := helper.APIResponse("success login user", 200, "success", gin.H{"token": token})
	c.JSON(200, response)
}
