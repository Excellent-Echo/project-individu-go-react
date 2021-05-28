package route

import (
	"project-individu-go-react/auth"
	"project-individu-go-react/config"
	"project-individu-go-react/layer"
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
	authService    = auth.NewService()
	userDelivery   = user.NewUserDeliver(userService, authService)

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

func UserRoutes(r *gin.Engine) {

	//User Route
	r.POST("/users/login", userDelivery.LoginUserDeliver)
	r.GET("/users", layer.Middleware(userService, authService), userDelivery.ShowUserDeliver)
	r.GET("/users/:user_id", layer.Middleware(userService, authService), userDelivery.GetUserByIDDeliver)
	r.POST("users/register", userDelivery.CreateUserDeliver)
	r.PUT("users/:user_id", layer.Middleware(userService, authService), userDelivery.UpdateUserByIDDeliver)
	r.DELETE("users/:user_id", layer.Middleware(userService, authService), userDelivery.DeleteUserByIDDeliver)

	//Customer Route
	r.GET("/customers", layer.Middleware(userService, authService), customerDelivery.ShowCustomerDeliver)
	r.GET("/customers/:cid", layer.Middleware(userService, authService), customerDelivery.GetCustomerByCIDDeliver)
	r.POST("customers/register", layer.Middleware(userService, authService), customerDelivery.CreateCustomerDeliver)
	r.PUT("/customers/:cid", layer.Middleware(userService, authService), customerDelivery.UpdateCustomerByCIDDeliver)
	r.DELETE("/customers/:cid", layer.Middleware(userService, authService), customerDelivery.DeleteCustomerByCIDDeliver)

	//Account Route
	r.GET("/accounts", layer.Middleware(userService, authService), accountDelivery.ShowAccountDeliver)
	r.GET("/accounts/:sid", layer.Middleware(userService, authService), accountDelivery.GetAccountBySIDDeliver)
	r.POST("/accounts/register", layer.Middleware(userService, authService), accountDelivery.CreateAccountDeliver)
	r.PUT("/accounts/:sid", layer.Middleware(userService, authService), accountDelivery.UpdateAccountBYSIDDeliver)
	r.DELETE("/accounts/:sid", layer.Middleware(userService, authService), accountDelivery.DeleteAccountBySIDDeliver)

	//History Route
	r.GET("/histories", layer.Middleware(userService, authService), historyDelivery.ShowHistoriesDeliver)
	r.GET("/history/:hid", layer.Middleware(userService, authService), historyDelivery.GetHistoryByHIDDeliver)
	r.GET("/histories/account/:sid", layer.Middleware(userService, authService), historyDelivery.GetHistoriesBySIDDeliver)
	r.POST("/history/new", layer.Middleware(userService, authService), historyDelivery.CreateHistoryDeliver)

}
