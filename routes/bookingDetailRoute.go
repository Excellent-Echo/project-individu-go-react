package routes

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/handler"
	"project-individu-go-react/layer/detailbooking"
)

var (
	bookingDetailRepository = detailbooking.NewRepository(DB)
	bookingDetailService    = detailbooking.NewService(bookingDetailRepository)
	bookingDetailHandler    = handler.NewBookingDetailHander(bookingDetailService, authService)
)

func BookingDetailRoute(router *gin.Engine) {
	router.GET("/booking-detail", handler.Middleware(userService, authService), bookingDetailHandler.ShowBookingDetailHandler)
	router.POST("/booking-detail", bookingDetailHandler.CreateBookingDetailHandler)
	router.GET("/booking-detail/:booking_detail_id", handler.Middleware(userService, authService), bookingDetailHandler.GetBookingDetailByIDHandler)
}
