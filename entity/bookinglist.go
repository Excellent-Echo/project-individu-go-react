package entity

import "time"

type BookingList struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
	User_id   []User    `gorm:"foreignKey:BookingListID" json:"userID"`
}
