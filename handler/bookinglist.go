package handler

import (
	"projectpenyewaanlapangan/auth"
	"projectpenyewaanlapangan/bookinglist"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookingHandler struct {
	bookService bookinglist.Service
	authService auth.Service
}

func NewBookingHandler(bookService bookinglist.Service, authService auth.Service) *bookingHandler {
	return &bookingHandler{bookService, authService}
}

func (h *bookingHandler) GetAllBookingHandler(c *gin.Context) {
	books, err := h.bookService.GetAllBookingList()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all todo", 200, "success", books)
	c.JSON(200, response)
}

func (h *bookingHandler) GetAllBookingbyUserIDAndFieldHandler(c *gin.Context) {

	userData := int(c.MustGet("currentUser").(int))
	fieldData := c.Params.ByName("field_id")

	userID := strconv.Itoa(userData)

	books, err := h.bookService.GetAllBookingbyUserIDAndField(userID, fieldData)

	if err != nil {
		responseError := helper.APIResponse("status unauthorize", 401, "error", gin.H{"error": err.Error()})

		c.JSON(401, responseError)
		return
	}

	response := helper.APIResponse("success get all todo by user ID", 200, "success", books)
	c.JSON(200, response)
}

func (h *bookingHandler) SaveNewBookingHandler(c *gin.Context) {
	userData := int(c.MustGet("currentUser").(int))
	fieldData := c.Params.ByName("field_id")

	if userData == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}

	userID := strconv.Itoa(userData)

	var inputBooking entity.BookingListInput

	if err := c.ShouldBindJSON(&inputBooking); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newBooking, err := h.bookService.SaveNewBooking(inputBooking, userID, fieldData)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new todo", 201, "success", newBooking)
	c.JSON(201, response)
}

func (h *bookingHandler) UpdateBookByIDHandler(c *gin.Context) {
	bookID := c.Params.ByName("booking_id")

	var updateBookingList entity.UpdateBookingListInput

	if err := c.ShouldBindJSON(&updateBookingList); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	book, err := h.bookService.GetBookingByID(bookID)

	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userData := int(c.MustGet("currentUser").(int))

	if book.UserID != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	books, err := h.bookService.UpdateBookByID(bookID, updateBookingList)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update todo by ID", 200, "success", books)
	c.JSON(200, response)
}
