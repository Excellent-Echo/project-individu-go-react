package config

import (
	"project-individu-go-react/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:marwanajunolii@tcp(localhost)/forum_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Post{})
	db.AutoMigrate(&entity.Category{})
	db.AutoMigrate(&entity.Comment{})
	db.AutoMigrate(&entity.Like{})

	return db
}
