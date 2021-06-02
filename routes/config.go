package routes

import (
	"projectpenyewaanlapangan/config"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.Connect()
)
