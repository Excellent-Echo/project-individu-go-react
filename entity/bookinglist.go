package entity

import "time"

type BookingList struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Date        time.Time `json:"date"`
	TimeForPlay int       `json:"time_for_play"`
	FieldLists  []FieldList
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// DeletedAt time.Time `gorm:"index" json:"-"`
}

type BookingListInput struct {
	Date        time.Time `json:"date" binding:"required"`
	TimeForPlay int       `json:"time_for_play" binding:"required"`
}

type UpdateBookingListInput struct {
	Date        time.Time `json:"date"`
	TimeForPlay int       `json:"time_for_play"`
}
