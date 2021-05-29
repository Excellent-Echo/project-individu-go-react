package entity

type BookingDetail struct {
	ID         int `gorm:"primaryKey" json:"id"`
	BookingID  int `json:"booking_id"`
	PsikologID int `json:"psikolog_id"`
}

type BookingDetailInput struct {
	BookingID  int `json:"booking_id" binding:"required"`
	PsikologID int `json:"psikolog_id" binding:"required"`
}
