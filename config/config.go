package config

import (
	"project-go-react/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	dsn := "root:@tcp(localhost)/database_toko?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.Category{})
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Order{})
	db.AutoMigrate(&entity.Product{})

	return db
}
