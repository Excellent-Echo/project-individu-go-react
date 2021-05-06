package main

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/handler"
)

//var DB = config.ConnectToDatabase()

func main() {
	//var booking []entity.Booking
	//
	//if err := DB.Find(&booking).Error ; err != nil {
	//	panic(err.Error())
	//}
	//
	//fmt.Println(booking)

	r := gin.Default()

	// Endpoint users
	r.GET("/users", handler.GetAllUser)
	r.POST("/users", handler.CreateNewUser)
	r.GET("users/:user_id", handler.GetUserByID)
	r.PUT("users/:user_id", handler.UpdateUserByID)
	r.DELETE("users/:user_id", handler.DeleteByUserID)

	// Endpoint role
	r.GET("/roles", handler.GetAllRole)
	r.POST("/roles", handler.CreateNewRole)
	r.GET("roles/:role_id", handler.GetRoleByID)
	r.PUT("roles/:role_id", handler.UpdateRoleByID)
	r.DELETE("roles/:role_id", handler.DeleteByRoleID)

	// Endpoint booking
	r.GET("/booking", handler.GetAllBooking)
	r.POST("/booking", handler.CreateNewBooking)
	r.GET("booking/:booking_id", handler.GetBookingByID)
	r.PUT("booking/:booking_id", handler.UpdateBookingByID)
	r.DELETE("booking/:booking_id", handler.DeleteByBookingID)

	r.Run(":3000")
}
