package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-individu-go-react/entity"
	"time"
)

func CreateNewPsikolog(c *gin.Context) {
	var getPsikolog entity.PsikologInput

	if err := c.ShouldBindJSON(&getPsikolog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	var newPsikolog = entity.Psikologi{
		Firstname:       getPsikolog.Firstname,
		Lastname:        getPsikolog.Lastname,
		Title:           getPsikolog.Title,
		Price:           getPsikolog.Price,
		JenisKonsultasi: getPsikolog.JenisKonsultasi,
		Description:     getPsikolog.Description,
		Review:          getPsikolog.Review,
		CreateAt:        time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := DB.Create(&newPsikolog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newPsikolog)

}

func GetAllPsikolog(c *gin.Context) {
	var psikolog []entity.Psikologi

	if err := DB.Find(&psikolog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, psikolog)
}

func GetPsikologByID(c *gin.Context) {
	var GetPsikologID entity.Psikologi

	id := c.Param("psikolog_id")

	if err := DB.Where("id = ?", id).Preload("Bookings").Preload("BookingDetails").Find(&GetPsikologID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	if GetPsikologID.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	c.JSON(http.StatusOK, GetPsikologID)
}

func UpdatePsikologByID(c *gin.Context) {
	var updatePsikolog entity.Psikologi
	var updatePsikologInput entity.PsikologInput

	id := c.Params.ByName("psikolog_id")

	if err := DB.Where("id = ?", id).Find(&updatePsikolog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	if updatePsikolog.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := c.ShouldBindJSON(&updatePsikologInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	updatePsikolog.Firstname = updatePsikologInput.Firstname
	updatePsikolog.Lastname = updatePsikologInput.Lastname
	updatePsikolog.Title = updatePsikologInput.Title
	updatePsikolog.Price = updatePsikologInput.Price
	updatePsikolog.JenisKonsultasi = updatePsikologInput.JenisKonsultasi
	updatePsikolog.Description = updatePsikologInput.Description
	updatePsikolog.Review = updatePsikologInput.Review
	updatePsikolog.UpdatedAt = time.Now()

	if err := DB.Where("id = ?", id).Save(&updatePsikolog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatePsikolog)
}

func DeleteByPsikologID(c *gin.Context) {
	id := c.Param("psikolog_id")

	var DeletePsikolog entity.Psikologi

	DB.Where("id = ?", id).Find(&DeletePsikolog)

	if DeletePsikolog.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := DB.Where("id = ?", id).Delete(&entity.Psikologi{}).Error; err != nil {
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
