package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `gorm:"primaryKey" json:"id`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
