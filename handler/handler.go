package handler

import (
	"log"
	"net/http"
	"project-individu-go-react/config"
	"project-individu-go-react/entity"

	"github.com/gin-gonic/gin"
)

var DB = config.Conn()

func GetAllKontak(c *gin.Context) {
	var kontak []entity.Kontak

	if err := DB.find(&kontak).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in internal server",
		})
		return
	}
	c.JSON(http.StatusOK, kontak)
}

func CreateNewKontak(c *gin.Context) {
	var getkontak entity.KontakInput

	if err := c.ShouldBindJSON(&getkontak); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"error_message": err.Error(),
		})
		return
	}

	var newKontak = entity.Kontak{
		NmKontak:   getkontak.NmKontak,
		NoTelp:     getkontak.NoTelp,
		Perusahaan: getkontak.Perusahaan,
		Email:      getkontak.Email,
	}

	if err := DB.create(&newKontak).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, newKontak)
}

func HandleIDKontak(c *gin.Context) {
	var kontak entity.Kontak

	id := c.Params.ByName("idKontak")

	if err := DB.Where("idKontak = ?", id).Find(&kontak).Error; err != nil {
		log.Println(err.Error)
	}

	c.JSON(http.StatusOK, kontak)
}

func HandleDelKontak(c *gin.Context) {
	id := c.Params.ByName("idKontak")

	if err := DB.Where("idKontak = ?", id).Delete(&entity.Kontak{}).Error; err != nil {
		log.Println(err.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"idKontak": id,
		"message":  "success delete",
	})
}

func HandleUpKontak(c *gin.Context) {
	var kontak entity.Kontak
	var kontakInput entity.KontakInput

	id := c.Params.ByName("idKontak")
	DB.Where("idKontak = ?", id).Find(&kontak)

	if err := c.ShouldBindJSON(&kontakInput); err != nil {
		log.Println(err.Error())
		return
	}

	kontak.NmKontak = kontakInput.NmKontak
	kontak.NoTelp = kontakInput.NoTelp
	kontak.Perusahaan = kontakInput.Perusahaan
	kontak.Email = kontakInput.Email

	if err := DB.Where("idKontak = ?", id).Save(&kontak).Error; err != nil {
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, kontak)
}
