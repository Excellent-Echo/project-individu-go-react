package entity

import "time"

type SportList struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	SportName string    `json:"sport_name"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type SportListInput struct {
	SportName string `json:"sport_name" binding:"required"`
}