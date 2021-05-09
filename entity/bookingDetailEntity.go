package entity

// BookingDetail struct tabel bookingdetail
type BookingDetail struct {
	ID         int `gorm:"primaryKey"`
	BookingID  int `json:"booking_id"`
	PsikologID int `json:"psikolog_id"`
}
