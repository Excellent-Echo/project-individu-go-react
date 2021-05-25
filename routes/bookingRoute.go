package routes

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/handler"
)

func BookingRoute(router *gin.Engine) {
	router.GET("/booking", handler.GetAllBooking)
	router.POST("/booking", handler.CreateNewBooking)
	router.GET("booking/:booking_id", handler.GetBookingByID)
	router.PUT("booking/:booking_id", handler.UpdateBookingByID)
	router.DELETE("booking/:booking_id", handler.DeleteByBookingID)
}
