package config

import (
	"fmt"
	"os"
	"project-individu-go-react/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	err := godotenv.Load()

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	// dsn := "root:marwanajunolii@tcp(localhost)/forum_app?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "XH5nOj1OqJ:7X580KSfLf@tcp(remotemysql.com:3306)/XH5nOj1OqJ?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Questions{})
	db.AutoMigrate(&entity.Tags{})
	db.AutoMigrate(&entity.Answers{})

	return db
}
