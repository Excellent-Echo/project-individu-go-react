package handler

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/auth"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/layer/booking"
)

type bookingHandler struct {
	bookingService booking.Service
	authService    auth.Service
}

func NewBookingHander(bookingService booking.Service, authService auth.Service) *bookingHandler {
	return &bookingHandler{bookingService, authService}
}

func (h *bookingHandler) ShowBookingHandler(c *gin.Context) {
	bookings, err := h.bookingService.GetAllBooking()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error internal server",
			"error":   err.Error(),
		})
		return
	}

	response := helper.APIResponse(
		"Success Get All Bookings Data",
		200,
		"Status OK",
		bookings,
	)

	c.JSON(200, response)
}

func (h *bookingHandler) CreateBookingHandler(c *gin.Context) {
	var NewBookingInput entity.BookingInput

	if err := c.ShouldBindJSON(&NewBookingInput); err != nil {
		splitError := helper.SplitErrorInformation(err)

		responseError := helper.APIResponse(
			"Input data required",
			400,
			"bad request",
			gin.H{
				"error": splitError,
			},
		)

		c.JSON(400, responseError)
		return
	}

	newBooking, err := h.bookingService.SaveNewBooking(NewBookingInput)

	if err != nil {
		responseError := helper.APIResponse(
			"Internal server error",
			500,
			"error",
			gin.H{
				"error": err.Error(),
			},
		)

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse(
		"Success Create new Bookings Data",
		201,
		"Status Created",
		newBooking,
	)

	c.JSON(201, response)
}

func (h *bookingHandler) GetBookingByIDHandler(c *gin.Context) {
	id := c.Params.ByName("booking_id")

	bookingByID, err := h.bookingService.GetBookingByID(id)

	if err != nil {
		responseError := helper.APIResponse("Error bad request get booking id", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("Success get booking by id", 200, "success", bookingByID)
	c.JSON(200, response)
}

func (h *bookingHandler) GetandDeleteBookingByIDHandler(c *gin.Context) {
	id := c.Params.ByName("booking_id")

	bookingDelete, err := h.bookingService.GetandDeleteBookingByID(id)

	if err != nil {
		responseError := helper.APIResponse("Error delete booking id", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	responseSuccess := helper.APIResponse("Success delete booking id", 200, "Delete OK", bookingDelete)
	c.JSON(200, responseSuccess)
}
