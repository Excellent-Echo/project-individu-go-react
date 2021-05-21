package config

import (
	"PROJECT-INDIVIDU-GO-REACT/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:bootcamp@tcp(localhost)/m021?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Mentor{})
	db.AutoMigrate(&entity.Booking{})
	return db
}
