package customer

import (
	"project-individu-go-react/entities"
	"project-individu-go-react/helper"

	"github.com/gin-gonic/gin"
)

type customerDeliver struct {
	customerService Service
}

func NewCustomerDeliver(customerService Service) *customerDeliver {
	return &customerDeliver{customerService}
}

func (d *customerDeliver) ShowCustomerDeliver(c *gin.Context) {
	customers, err := d.customerService.GetAllCustomer()

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responError)

		return
	}

	response := helper.APIResponse("success get all customer", 200, "status OK", customers)
	c.JSON(200, response)
}

func (d *customerDeliver) GetCustomerByCIDDeliver(c *gin.Context) {
	cid := c.Params.ByName("cid")

	customer, err := d.customerService.GetCustomerByCID(cid)

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responError)

		return
	}

	response := helper.APIResponse("success get customer by cid", 200, "status OK", customer)
	c.JSON(200, response)
}

func (d *customerDeliver) CreateCustomerDeliver(c *gin.Context) {
	var inputCustomer entities.CostumerInput

	if err := c.ShouldBindJSON(&inputCustomer); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("Input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newCustomer, err := d.customerService.SaveNewCustomer(inputCustomer)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
	}

	response := helper.APIResponse("success create new customer", 201, "status Created", newCustomer)

	c.JSON(201, response)
}

func (d *customerDeliver) UpdateCustomerByCIDDeliver(c *gin.Context) {
	cid := c.Params.ByName("cid")

	var updateCustomerInput entities.CostumerUpdateInput

	if err := c.ShouldBindJSON(&updateCustomerInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("Input data required", 400, "bad request", gin.H{"error": splitError})

		c.JSON(400, responseError)
		return
	}

	customer, err := d.customerService.UpdateCustomerByCID(cid, updateCustomerInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("succes update customer by cid", 200, "status OK", customer)
	c.JSON(200, response)
}

func (d *customerDeliver) DeleteCustomerByCIDDeliver(c *gin.Context) {
	cid := c.Params.ByName("cid")

	customer, err := d.customerService.DeleteCustomerByCID(cid)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete user", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success delete customer by cid", 201, "status OK", customer)
	c.JSON(200, response)
}
