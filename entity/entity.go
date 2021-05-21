package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID          int `gorm:"primaryKey"`
	Address     string
	Gender      string
	NoHandphone int
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Course struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	CategoryID  int
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	// Courses     []Course `gorm:"foreignKey:CategoryID"`
}
