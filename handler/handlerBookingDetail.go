package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-individu-go-react/entity"
)

func CreateNewBookingDetail(c *gin.Context) {
	var getBookingDetail entity.BookingDetail

	if err := c.ShouldBindJSON(&getBookingDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	var newBookingDetail = entity.BookingDetail{
		BookingID:  getBookingDetail.BookingID,
		PsikologID: getBookingDetail.PsikologID,
	}

	if err := DB.Create(&newBookingDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newBookingDetail)

}

func GetAllBookingDetail(c *gin.Context) {
	var bookingDetail []entity.BookingDetail

	if err := DB.Find(&bookingDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, bookingDetail)
}

func GetBookingDetailByID(c *gin.Context) {
	var GetBookingDetailID entity.BookingDetail

	id := c.Param("booking_detail_id")

	if err := DB.Where("id = ?", id).Find(&GetBookingDetailID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	if GetBookingDetailID.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	c.JSON(http.StatusOK, GetBookingDetailID)
}

func UpdateBookingDetailByID(c *gin.Context) {
	var updateBookingDetail entity.BookingDetail

	id := c.Params.ByName("booking_detail_id")

	if err := DB.Where("id = ?", id).Find(&updateBookingDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	if updateBookingDetail.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := c.ShouldBindJSON(&updateBookingDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	var newBookingDetail = entity.BookingDetail{
		BookingID:  updateBookingDetail.BookingID,
		PsikologID: updateBookingDetail.PsikologID,
	}

	if err := DB.Where("id = ?", id).Save(&newBookingDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newBookingDetail)
}

func DeleteByBookingDetailID(c *gin.Context) {
	id := c.Param("booking_detail_id")

	var DeleteBookingDetail entity.BookingDetail

	DB.Where("id = ?", id).Find(&DeleteBookingDetail)

	if DeleteBookingDetail.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := DB.Where("id = ?", id).Delete(&entity.BookingDetail{}).Error; err != nil {
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
