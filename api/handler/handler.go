package handler

// import (
// 	"doctorsAppointment/config"
// 	"doctorsAppointment/entity"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// var DB = config.Connection()

// func GetAllDoctor(c *gin.Context) {
// 	var doctors []entity.Doctor

// 	if err := DB.Preload("MedicalSpecialist").Preload("Appointment").Find(&doctors).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":        "error in internal server",
// 			"message_error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, doctors)
// }

// func GetAllPatient(c *gin.Context) {
// 	var patients []entity.Patient

// 	if err := DB.Preload("Appointment").Find(&patients).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":        "error in internal server",
// 			"message_error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, patients)
// }

// func GetAllAppointment(c *gin.Context) {
// 	var appointments []entity.Appointment

// 	if err := DB.Find(&appointments).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":        "error in internal server",
// 			"message_error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, appointments)
// }
