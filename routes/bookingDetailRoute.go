package routes

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/handler"
)

func BookingDetailRoute(router *gin.Engine) {
	router.GET("/booking-detail", handler.GetAllBookingDetail)
	router.POST("/booking-detail", handler.CreateNewBookingDetail)
	router.GET("booking-detail/:booking_detail_id", handler.GetBookingDetailByID)
	router.PUT("booking-detail/:booking-detail_id", handler.UpdateBookingDetailByID)
	router.DELETE("booking-detail/:booking_detail_id", handler.DeleteByBookingDetailID)
}
