package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project-individu-go-react/entity"
)

func ConnectToDatabase() *gorm.DB {
	dsn := "root:root@tcp(localhost)/konsultasi_psikolog?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Role{})
	db.AutoMigrate(&entity.Booking{})
	db.AutoMigrate(&entity.BookingDetail{})
	db.AutoMigrate(&entity.Psikologi{})

	return db
}
