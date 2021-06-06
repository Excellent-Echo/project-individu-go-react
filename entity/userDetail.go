package entity

import (
	"time"
)

type UserDetail struct {
	ID          int       `gorm:"primaryKey" json:"udetail_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	NoHandphone int       `json:"no_handphone"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserDetailInput struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	NoHandphone uint   `json:"no_handphone" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type UpdateUserDetailInput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	NoHandphone int    `json:"no_handphone"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
}
