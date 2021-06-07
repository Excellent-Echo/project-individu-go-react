package config

import (
	"project-individu-go-react/migration"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// dbUser := os.Getenv("DB_USERNAME")
	// dbPass := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	dsn := "root:@tcp(localhost)/coding_here"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&migration.User{})
	db.AutoMigrate(&migration.Category{})
	db.AutoMigrate(&migration.UserDetail{})
	db.AutoMigrate(&migration.Course{})

	return db

}
