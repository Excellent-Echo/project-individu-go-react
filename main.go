package main

import (
	"project-individu-go-react/config"
	"project-individu-go-react/customer"
	"project-individu-go-react/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.Connect()

	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)
	userDelivery   = user.NewUserDeliver(userService)

	customerRepository = customer.NewRepository(DB)
	customerService    = customer.NewService(customerRepository)
	customerDelivery   = customer.NewCustomerDeliver(customerService)
)

func main() {
	r := gin.Default()

	//User Route
	r.GET("/users", userDelivery.ShowUserDeliver)
	r.GET("/users/:user_id", userDelivery.GetUserByIDDeliver)
	r.POST("users/register", userDelivery.CreateUserDeliver)
	r.PUT("users/:user_id", userDelivery.UpdateUserByIDDeliver)
	r.DELETE("users/:user_id", userDelivery.DeleteUserByIDDeliver)

	//Customer Route
	r.GET("/customers", customerDelivery.ShowCustomerDeliver)
	r.GET("/customers/:cid", customerDelivery.GetCustomerByCIDDeliver)
	r.POST("customers/register", customerDelivery.CreateCustomerDeliver)
	r.PUT("/customers/:cid", customerDelivery.UpdateCustomerByCIDDeliver)
	r.DELETE("/customers/:cid", customerDelivery.DeleteCustomerByCIDDeliver)

	r.Run(":8888")
}
