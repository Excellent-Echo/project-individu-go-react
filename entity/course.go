package entity

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID          int            `gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CategoryID  int            `json:"category_id"`
	UserID      int            `json:"user_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
