package entity

import "time"

type BookingList struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt time.Time `gorm:"index" json:"-"`
	User_id   int    `json:"userID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type BookingListInput struct {
	Date      time.Time `json:"date"`
	User_id   int       `json:"userID"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}
