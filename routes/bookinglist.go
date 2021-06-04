package routes

import (
	"projectpenyewaanlapangan/bookinglist"
	"projectpenyewaanlapangan/handler"

	"github.com/gin-gonic/gin"
)

var (
	bookingRepository = bookinglist.NewRepository(DB)
	bookingService    = bookinglist.NewService(bookingRepository)
	bookingHandler    = handler.NewBookingHandler(bookingService, authService)
)

func BookingLIstRoute(r *gin.Engine) {
	r.GET("/bookinglist", handler.Middleware(userService, authService), bookingHandler.GetAllBookingHandler)
	r.POST("/bookinglist", handler.Middleware(userService, authService), bookingHandler.SaveNewBookingHandler)
	r.PUT("/bookinglist/:booking_id", handler.Middleware(userService, authService), bookingHandler.UpdateBookByIDHandler)
}
