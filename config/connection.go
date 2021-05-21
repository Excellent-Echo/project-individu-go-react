package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "Radika:!Satu234limA!@tcp(localhost)/watchlist?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	//db.AutoMigrate(&entities.User{})
	//db.AutoMigrate(&entities.UserDetail{})
	//db.AutoMigrate(&entities.Watchlist{})
	return db
}
