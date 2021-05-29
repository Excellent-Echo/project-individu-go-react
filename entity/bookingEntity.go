package entity

import "time"

type Booking struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      int       `json:"user_id"`
	PsikologID  int       `json:"psikolog_id"`
	BookingDate int       `json:"booking_date"`
	BookingTime int       `json:"booking_time"`
	CreateAt    time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	//DeletedAt      time.Time       `gorm:"index" json:"-"`
	BookingDetails []BookingDetail `gorm:"foreignKey:BookingID"`
}

type BookingInput struct {
	UserID      int `json:"user_id" binding:"required"`
	PsikologID  int `json:"psikolog_id" binding:"required"`
	BookingDate int `json:"booking_date" binding:"required"`
	BookingTime int `json:"booking_time" binding:"required"`
}
