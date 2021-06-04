package entity

import "time"

type FieldList struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	FieldName     string `json:"field_name"`
	FieldImage    string `json:"field_image"`
	RentPrice     int    `json:"rent_price"`
	BookingListID int    `json:"booking_list_id"`
	Users         []User
	CreatedAt     time.Time `json:"create_at"`
	UpdatedAt     time.Time `json:"update_at"`
	// DeletedAt  time.Time `gorm:"index" json:"-"`
}

type FieldListInput struct {
	FieldName     string `json:"field_name" binding:"required"`
	FieldImage    string `json:"field_image" binding:"required"`
	RentPrice     int    `json:"rent_price" binding:"required"`
	BookingListID int    `json:"booking_list_id" binding:"required"`
}
