package handler

import (
	"project-individu-go-react/auth"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/layer/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.UserService
	authService auth.Service
}

func NewUserHandler(userService user.UserService, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
	id := c.Param("id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userResponse := helper.APIResponse("get user succeed", 200, "success", user)
	c.JSON(200, userResponse)
}

func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	var updateUserInput entity.UpdateUserInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	idParam, _ := strconv.Atoi(id)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	user, err := h.userService.UpdateUserByID(id, updateUserInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("update user succeed", 200, "success", user)
	c.JSON(200, response)
}

func (h *userHandler) DeleteByUserIDHandler(c *gin.Context) {
	id := c.Param("id")

	idParam, _ := strconv.Atoi(id)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	user, err := h.userService.DeleteByUserID(id)

	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userResponse := helper.APIResponse("user was deleted successfully", 200, "success", user)
	c.JSON(200, userResponse)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser entity.LoginUserInput

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseError := helper.APIResponse("input data error", 401, "bad request", gin.H{"errors": err})

		c.JSON(401, responseError)
		return
	}

	token, err := h.authService.GenerateToken(int(userData.ID))
	if err != nil {
		responseError := helper.APIResponse("input data error", 401, "bad request", gin.H{"errors": err})

		c.JSON(401, responseError)
		return
	}
	response := helper.APIResponse("login user succeed", 200, "success", gin.H{"token": token})
	c.JSON(200, response)
}