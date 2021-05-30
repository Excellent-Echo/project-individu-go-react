package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"project-individu-go-react/migration"
)

// ConnectToDatabase untuk koneksi dari database MySQL
func ConnectToDatabase() *gorm.DB {
	//err := godotenv.Load()

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	//dsn := "cb4RUILMib:WOiDCnT5Ey@tcp(remotemysql.com:3306)/cb4RUILMib?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// Untuk membuat tabel baru ke database MySQL
	db.AutoMigrate(&migration.User{})
	db.AutoMigrate(&migration.Role{})
	db.AutoMigrate(&migration.Booking{})
	db.AutoMigrate(&migration.BookingDetail{})
	db.AutoMigrate(&migration.Psikologi{})
	db.AutoMigrate(&migration.UserProfile{})

	return db
}
