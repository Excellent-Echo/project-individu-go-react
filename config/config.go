package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectToDatabase untuk koneksi dari database MySQL
func ConnectToDatabase() *gorm.DB {
	dsn := "root:root@tcp(localhost)/konsultasi_psikolog?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// Untuk membuat tabel baru ke database MySQL
	//db.AutoMigrate(&entity.User{})
	//db.AutoMigrate(&entity.Role{})
	//db.AutoMigrate(&entity.Booking{})
	//db.AutoMigrate(&entity.BookingDetail{})
	//db.AutoMigrate(&entity.Psikologi{})

	return db
}
