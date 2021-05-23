package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "Radika:!Satu234limA!@tcp(localhost)/AboutMoney?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// db.AutoMigrate(&entities.User{})
	// db.AutoMigrate(&entities.Costumer{})
	// db.AutoMigrate(&entities.Account{})
	// db.AutoMigrate(&entities.History{})
	return db
}
