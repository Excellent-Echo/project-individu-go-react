package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Address     string         `json:"address"`
	Gender      string         `json:"gender"`
	NoHandphone int            `json:"no_handphone"`
	UserID      int            `json:"user_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserDetailInput struct {
	Address     string `json:"address"`
	NoHandphone int    `json:"no_handphone" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
}

type UpdateUserDetailInput struct {
	Address     string `json:"address"`
	NoHandphone int    `json:"no_handphone"`
	Gender      string `json:"gender"`
}
