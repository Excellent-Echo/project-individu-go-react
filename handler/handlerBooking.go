package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-individu-go-react/entity"
	"time"
)

func CreateNewBooking(c *gin.Context) {
	var getBooking entity.BookingInput

	if err := c.ShouldBindJSON(&getBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	var newBooking = entity.Booking{
		UserID:      getBooking.UserID,
		PsikologID:  getBooking.PsikologID,
		BookingTime: getBooking.BookingTime,
		BookingDate: getBooking.BookingDate,
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := DB.Create(&newBooking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newBooking)

}

func GetAllBooking(c *gin.Context) {
	var booking []entity.Booking

	if err := DB.Find(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func GetBookingByID(c *gin.Context) {
	var GetBookingID entity.Booking

	id := c.Param("booking_id")

	if err := DB.Where("id = ?", id).Preload("BookingDetails").Find(&GetBookingID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	if GetBookingID.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	c.JSON(http.StatusOK, GetBookingID)
}

func UpdateBookingByID(c *gin.Context) {
	var updateBooking entity.Booking
	var updateBookingInput entity.BookingInput

	id := c.Params.ByName("booking_id")

	if err := DB.Where("id = ?", id).Find(&updateBooking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	if updateBooking.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := c.ShouldBindJSON(&updateBookingInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	updateBooking.BookingDate = updateBookingInput.BookingDate
	updateBooking.BookingTime = updateBookingInput.BookingTime
	updateBooking.UpdatedAt = time.Now()

	if err := DB.Where("id = ?", id).Save(&updateBooking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updateBooking)
}

func DeleteByBookingID(c *gin.Context) {
	id := c.Param("booking_id")

	var DeleteBooking entity.Booking

	DB.Where("id = ?", id).Find(&DeleteBooking)

	if DeleteBooking.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := DB.Where("id = ?", id).Delete(&entity.Booking{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "delete success",
		"message": "user id " + id + " successfull delete",
	})
}
