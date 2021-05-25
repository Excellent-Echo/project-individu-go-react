package entity

import "time"

// Booking struct tabel bookings
type Booking struct {
	ID          int       `gorm:"primaryKey"`
	UserID      int       `json:"user_id"`
	PsikologID  int       `json:"psikolog_id"`
	BookingDate int       `json:"booking_date"`
	BookingTime int       `json:"booking_time"`
	CreateAt    time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	//DeletedAt      time.Time       `gorm:"index" json:"-"`
	//BookingDetails []BookingDetail `gorm:"foreignKey:BookingID"`
}

type BookingInput struct {
	UserID      int `json:"user_id"`
	PsikologID  int `json:"psikolog_id"`
	BookingDate int `json:"booking_date"`
	BookingTime int `json:"booking_time"`
}
