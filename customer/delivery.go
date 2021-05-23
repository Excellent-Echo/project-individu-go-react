package customer

import (
	"project-individu-go-react/helper"

	"github.com/gin-gonic/gin"
)

type customerDeliver struct {
	customerService Service
}

func NewUserDeliver(customerService Service) *customerDeliver {
	return &customerDeliver{customerService}
}

func (d *customerDeliver) ShowCustomerDeliver(c *gin.Context) {
	customers, err := d.customerService.GetAllCustomer()

	if err != nil {
		responError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responError)

		return
	}

	response := helper.APIResponse("success get all user", 200, "status OK", customers)
	c.JSON(200, response)
}
