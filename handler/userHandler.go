package handler

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHander(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// ShowUserHandler untuk mengambil semua data user, jika user mengakses dari endpoint "/users"
func (h *userHandler) ShowUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error internal server",
			"error":   err.Error(),
		})
		return
	}

	response := helper.APIResponse(
		"Success Get All Users Data",
		200,
		"Status OK",
		users,
	)

	c.JSON(200, response)
}

// CreateUserHandler untuk membuat data user baru, jika user menginput dari endpoint "/users/register"
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var NewUserInput entity.UserInput

	if err := c.ShouldBindJSON(&NewUserInput); err != nil {
		splitError := helper.SplitErrorInformation(err)

		responseError := helper.APIResponse(
			"Input data required",
			400,
			"bad request",
			gin.H{
				"error": splitError,
			},
		)

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(NewUserInput)

	if err != nil {
		responseError := helper.APIResponse(
			"Internal server error",
			500,
			"error",
			gin.H{
				"error": err.Error(),
			},
		)

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse(
		"Success Create new Users Data",
		201,
		"Status Created",
		newUser,
	)

	c.JSON(201, response)
}

// GetUserByIDHandler untuk mencari data user berdasarkan id nya, jika user menginput dari endpoint "/users/user_id"
func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	userByID, err := h.userService.GetUserByID(id)

	if err != nil {
		responseError := helper.APIResponse("Error bad request get user id", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("Success get user by id", 200, "success", userByID)
	c.JSON(200, response)
}

// GetandDeleteUserByIDHandler untuk menghapus data user berdasarkan id, jika user menginput dari endpoint "/users/user_id"
func (h *userHandler) GetandDeleteUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	userDelete, err := h.userService.GetandDeleteUserByID(id)

	if err != nil {
		responseError := helper.APIResponse("Error delete user id", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	responseSuccess := helper.APIResponse("Success delete user id", 200, "Delete OK", userDelete)
	c.JSON(200, responseSuccess)
}

// GetandUpdateUserByIDHandler untuk menghapus data user berdasarkan id, jika user menginput dari endpoint "/users/user_id"
func (h *userHandler) GetandUpdateUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var userUpdate entity.UserInputUpdate

	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": splitError})

		c.JSON(400, responseError)
		return
	}

	users, err := h.userService.GetandUpdateUserByID(id, userUpdate)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user by id", 200, "success", users)
	c.JSON(200, response)
}
