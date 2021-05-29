package handler

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/auth"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/layer/detailbooking"
)

type bookingDetailHandler struct {
	bookingDetailService detailbooking.Service
	authService          auth.Service
}

func NewBookingDetailHander(bookingDetailService detailbooking.Service, authService auth.Service) *bookingDetailHandler {
	return &bookingDetailHandler{bookingDetailService, authService}
}

func (h *bookingDetailHandler) ShowBookingDetailHandler(c *gin.Context) {
	bookingDetails, err := h.bookingDetailService.GetAllBookingDetail()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error internal server",
			"error":   err.Error(),
		})
		return
	}

	response := helper.APIResponse(
		"Success Get All Booking Details Data",
		200,
		"Status OK",
		bookingDetails,
	)

	c.JSON(200, response)
}

func (h *bookingDetailHandler) CreateBookingDetailHandler(c *gin.Context) {
	var NewBookingDetailInput entity.BookingDetailInput

	if err := c.ShouldBindJSON(&NewBookingDetailInput); err != nil {
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

	newBookingDetail, err := h.bookingDetailService.SaveNewBookingDetail(NewBookingDetailInput)

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
		"Success Create new Booking Detail Data",
		201,
		"Status Created",
		newBookingDetail,
	)

	c.JSON(201, response)
}

func (h *bookingDetailHandler) GetBookingDetailByIDHandler(c *gin.Context) {
	id := c.Params.ByName("booking_detail_id")

	bookingDetailByID, err := h.bookingDetailService.GetBookingDetailByID(id)

	if err != nil {
		responseError := helper.APIResponse("Error bad request get booking detail id", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("Success get booking detail by id", 200, "success", bookingDetailByID)
	c.JSON(200, response)
}
