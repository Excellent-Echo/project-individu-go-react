package entity

import "time"

type BookingList struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Date        time.Time `json:"date"`
	UserID      int       `json:"user_id"`
	FieldID     int       `json:"field_id"`
	TimeForPlay int       `json:"time_for_play"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// DeletedAt time.Time `gorm:"index" json:"-"`
}

type BookingListInput struct {
	Date        time.Time `json:"date"`
	UserID      int       `json:"user_id"`
	FieldID     int       `json:"field_id"`
	TimeForPlay int       `json:"time_for_play"`
}
