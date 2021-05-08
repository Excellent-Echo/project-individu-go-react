package main

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/config"
	"project-individu-go-react/handler"
	"project-individu-go-react/user"
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

	r := gin.Default()

	// Endpoint untuk model users
	r.GET("/users", userHandler.ShowUserHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.GET("users/:user_id", userHandler.GetUserByIDHandler)
	r.PUT("users/:user_id", handler.UpdateUserByID)
	r.DELETE("users/:user_id", userHandler.GetandDeleteUserByIDHandler)

	// Endpoint untuk model role
	r.GET("/roles", handler.GetAllRole)
	r.POST("/roles", handler.CreateNewRole)
	r.GET("roles/:role_id", handler.GetRoleByID)
	r.PUT("roles/:role_id", handler.UpdateRoleByID)
	r.DELETE("roles/:role_id", handler.DeleteByRoleID)

	// Endpoint untuk model booking
	r.GET("/booking", handler.GetAllBooking)
	r.POST("/booking", handler.CreateNewBooking)
	r.GET("booking/:booking_id", handler.GetBookingByID)
	r.PUT("booking/:booking_id", handler.UpdateBookingByID)
	r.DELETE("booking/:booking_id", handler.DeleteByBookingID)

	// Endpoint untuk model booking detail
	r.GET("/booking-detail", handler.GetAllBookingDetail)
	r.POST("/booking-detail", handler.CreateNewBookingDetail)
	r.GET("booking-detail/:booking_detail_id", handler.GetBookingDetailByID)
	r.PUT("booking-detail/:booking-detail_id", handler.UpdateBookingDetailByID)
	r.DELETE("booking-detail/:booking_detail_id", handler.DeleteByBookingDetailID)

	// Endpoint untuk model psikolog
	r.GET("/psikolog", handler.GetAllPsikolog)
	r.POST("/psikolog", handler.CreateNewPsikolog)
	r.GET("psikolog/:psikolog_id", handler.GetPsikologByID)
	r.PUT("psikolog/:psikolog_id", handler.UpdatePsikologByID)
	r.DELETE("psikolog/:psikolog_id", handler.DeleteByPsikologID)

	// Untuk menjalanakan server ke localhost komputer
	r.Run(":3000")

}
