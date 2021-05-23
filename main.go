package main

import (
	"project-individu-go-react/config"
	"project-individu-go-react/layer/account"
	"project-individu-go-react/layer/customer"
	"project-individu-go-react/layer/history"
	"project-individu-go-react/layer/user"

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

	accountRepository = account.NewRepository(DB)
	accountService    = account.NewService(accountRepository)
	accountDelivery   = account.NewAccountDeliver(accountService)

	historyRepository = history.NewRepository(DB)
	historyService    = history.NewService(historyRepository)
	historyDelivery   = history.NewHistoryDeliver(historyService)
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

	//Account Route
	r.GET("/accounts", accountDelivery.ShowAccountDeliver)
	r.GET("/accounts/:sid", accountDelivery.GetAccountBySIDDeliver)
	r.POST("/accounts/register", accountDelivery.CreateAccountDeliver)
	r.PUT("/accounts/:sid", accountDelivery.UpdateAccountBYSIDDeliver)
	r.DELETE("/accounts/:sid", accountDelivery.DeleteAccountBySIDDeliver)

	//History Route
	r.GET("/histories", historyDelivery.ShowHistoriesDeliver)
	r.GET("/history/:hid", historyDelivery.GetHistoryByHIDDeliver)
	r.GET("/histories/account/:sid", historyDelivery.GetHistoriesBySIDDeliver)
	r.POST("/history/new", historyDelivery.CreateHistoryDeliver)

	r.Run(":8888")
}
