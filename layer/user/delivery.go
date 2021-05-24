package user

import (
	"project-individu-go-react/auth"
	"project-individu-go-react/entities"
	"project-individu-go-react/helper"

	"github.com/gin-gonic/gin"
)

type userDeliver struct {
	userService Service
	authService auth.Service
}

func NewUserDeliver(userService Service, authService auth.Service) *userDeliver {
	return &userDeliver{userService, authService}
}

func (d *userDeliver) ShowUserDeliver(c *gin.Context) {
	users, err := d.userService.GetAllUser()

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responError)

		return
	}

	response := helper.APIResponse("success get all user", 200, "status OK", users)
	c.JSON(200, response)
}

func (d *userDeliver) GetUserByIDDeliver(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := d.userService.GetUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("error bad request user ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "succes", user)
	c.JSON(200, response)
}

func (d *userDeliver) CreateUserDeliver(c *gin.Context) {
	var inputUser entities.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("Input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newUser, err := d.userService.SaveNewUser(inputUser)

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responError)
		return
	}

	response := helper.APIResponse("success create new user", 201, "status Created", newUser)
	c.JSON(201, response)
}

func (d *userDeliver) UpdateUserByIDDeliver(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var updateUserInput entities.UpdateUserInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		//splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("Input data required", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	user, err := d.userService.UpdateUserByID(id, updateUserInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user by ID", 200, "succes", user)
	c.JSON(200, response)
}

func (d *userDeliver) DeleteUserByIDDeliver(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := d.userService.DeleteUserByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete user", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success delete user by ID", 200, "success", user)
	c.JSON(200, response)
}

//login delivery

func (d *userDeliver) LoginUserDeliver(c *gin.Context) {
	var inputLoginUser entities.LoginUserInput

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userData, err := d.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseError := helper.APIResponse("input data error", 401, "bad request", gin.H{"errors": err})

		c.JSON(401, responseError)
		return
	}

	token, err := d.authService.GenerateToken(userData.ID)
	if err != nil {
		responseError := helper.APIResponse("input data error", 401, "bad request", gin.H{"errors": err})

		c.JSON(401, responseError)
		return
	}

	response := helper.APIResponse("success login user", 200, "success", token)
	c.JSON(200, response)
}
