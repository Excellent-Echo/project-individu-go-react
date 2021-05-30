package handler

import (
	"github.com/gin-gonic/gin"
	"project-go-react/entity"
	"project-go-react/helper"
	"project-go-react/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

//show all user in database from route "/users"
func (h *userHandler) ShowUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("successget all user", 200, "status OK", users)
	c.JSON(200, response)
}

//user/external handler for create new user from "/user"
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)

		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(inputUser)

	if err != nil {

		responseError := helper.APIResponse("internal server error", 4, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("successget all user", 201, "status created", newUser)
	c.JSON(200, response)
}

func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Param("user_id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)

		return
	}

	response := helper.APIResponse("successget get user by id", 200, "succes", user)
	c.JSON(200, response)

}

func (h *userHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")
	user, err := h.userService.DeleteByUserID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)

		return
	}

	response := helper.APIResponse("success delete user by ID", 200, "succes", user)
	c.JSON(200, response)
}
