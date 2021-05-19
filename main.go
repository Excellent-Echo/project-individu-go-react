package main

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/config"
	"project-individu-go-react/handler"
	"project-individu-go-react/layer/user"
)

// Database, Repository, Service dan Handler
var (
	DB             = config.ConnectToDatabase()
	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)
	userHandler    = handler.NewUserHander(userService)
)

func main() {
	//var booking []entity.Booking
	//
	//if err := DB.Find(&booking).Error ; err != nil {
	//	panic(err.Error())
	//}
	//
	//fmt.Println(booking)

	router := gin.Default()

	// Endpoint untuk model users
	router.GET("/users", userHandler.ShowUserHandler)
	router.POST("/users/register", userHandler.CreateUserHandler)
	//router.POST("/users/login", userHandler.CreateUserHandler)
	router.GET("users/:user_id", userHandler.GetUserByIDHandler)
	router.PUT("users/:user_id", userHandler.GetandUpdateUserByIDHandler)
	//router.PUT("users/:user_id", handler.UpdateUserByID)
	router.DELETE("users/:user_id", userHandler.GetandDeleteUserByIDHandler)

	// Endpoint untuk model role
	router.GET("/roles", handler.GetAllRole)
	router.POST("/roles", handler.CreateNewRole)
	router.GET("roles/:role_id", handler.GetRoleByID)
	router.PUT("roles/:role_id", handler.UpdateRoleByID)
	router.DELETE("roles/:role_id", handler.DeleteByRoleID)

	// Endpoint untuk model booking
	router.GET("/booking", handler.GetAllBooking)
	router.POST("/booking", handler.CreateNewBooking)
	router.GET("booking/:booking_id", handler.GetBookingByID)
	router.PUT("booking/:booking_id", handler.UpdateBookingByID)
	router.DELETE("booking/:booking_id", handler.DeleteByBookingID)

	// Endpoint untuk model booking detail
	router.GET("/booking-detail", handler.GetAllBookingDetail)
	router.POST("/booking-detail", handler.CreateNewBookingDetail)
	router.GET("booking-detail/:booking_detail_id", handler.GetBookingDetailByID)
	router.PUT("booking-detail/:booking-detail_id", handler.UpdateBookingDetailByID)
	router.DELETE("booking-detail/:booking_detail_id", handler.DeleteByBookingDetailID)

	// Endpoint untuk model psikolog
	router.GET("/psikolog", handler.GetAllPsikolog)
	router.POST("/psikolog", handler.CreateNewPsikolog)
	router.GET("psikolog/:psikolog_id", handler.GetPsikologByID)
	router.PUT("psikolog/:psikolog_id", handler.UpdatePsikologByID)
	router.DELETE("psikolog/:psikolog_id", handler.DeleteByPsikologID)

	// Untuk menjalanakan server ke localhost komputer
	router.Run(":3000")

}
