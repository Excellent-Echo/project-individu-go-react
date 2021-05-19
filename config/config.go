package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// ConnectToDatabase untuk koneksi dari database MySQL
func ConnectToDatabase() *gorm.DB {
	err := godotenv.Load()

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

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
