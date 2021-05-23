package account

import (
	"project-individu-go-react/entities"
	"project-individu-go-react/helper"

	"github.com/gin-gonic/gin"
)

type accountDeliver struct {
	accountService Service
}

func NewAccountDeliver(accountService Service) *accountDeliver {
	return &accountDeliver{accountService}
}

func (d *accountDeliver) ShowAccountDeliver(c *gin.Context) {
	accounts, err := d.accountService.GetAllAccount()

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responError)

		return
	}

	response := helper.APIResponse("success get all account", 200, "status OK", accounts)
	c.JSON(200, response)
}

func (d *accountDeliver) GetAccountBySIDDeliver(c *gin.Context) {
	sid := c.Params.ByName("sid")

	account, err := d.accountService.GetAccountBySID(sid)

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responError)

		return
	}

	response := helper.APIResponse("success get account by sid", 200, "status OK", account)
	c.JSON(200, response)
}

func (d *accountDeliver) CreateAccountDeliver(c *gin.Context) {
	var inputAccount entities.AccountInput

	if err := c.ShouldBindJSON(&inputAccount); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("Input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newAccount, err := d.accountService.SaveNewAccount(inputAccount)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
	}

	response := helper.APIResponse("success create new account", 201, "status Created", newAccount)

	c.JSON(201, response)
}

func (d *accountDeliver) UpdateAccountBYSIDDeliver(c *gin.Context) {
	sid := c.Params.ByName("sid")

	var updateAccountInput entities.AccountUpdateInput

	if err := c.ShouldBindJSON(&updateAccountInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("Input data required", 400, "bad request", gin.H{"error": splitError})

		c.JSON(400, responseError)
		return
	}

	account, err := d.accountService.UpdateAccountBySID(sid, updateAccountInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("succes update account by sid", 200, "status OK", account)
	c.JSON(200, response)
}

func (d *accountDeliver) DeleteAccountBySIDDeliver(c *gin.Context) {
	sid := c.Params.ByName("sid")

	account, err := d.accountService.DeleteAccountBySID(sid)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete user", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success delete account by sid", 201, "status OK", account)
	c.JSON(200, response)
}
