package routes

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/handler"
	"project-individu-go-react/layer/booking"
)

var (
	bookingRepository = booking.NewRepository(DB)
	bookingService    = booking.NewService(bookingRepository)
	bookingHandler    = handler.NewBookingHander(bookingService, authService)
)

func BookingRoute(route *gin.Engine) {
	route.GET("/booking", handler.Middleware(userService, authService), bookingHandler.ShowBookingHandler)
	route.POST("/booking/order", bookingHandler.CreateBookingHandler)
	route.GET("booking/:booking_id", handler.Middleware(userService, authService), bookingHandler.GetBookingByIDHandler)
	route.DELETE("booking/:booking_id", handler.Middleware(userService, authService), bookingHandler.GetandDeleteBookingByIDHandler)

}
