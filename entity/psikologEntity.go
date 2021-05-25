package entity

import "time"

// Psikologi struct tabel psikolog
type Psikologi struct {
	ID              int       `gorm:"primaryKey" json:"id"`
	Firstname       string    `json:"firstname"`
	Lastname        string    `json:"lastname"`
	Title           string    `json:"title"`
	Price           int       `json:"price"`
	JenisKonsultasi string    `json:"jenis_konsultasi"`
	Description     string    `json:"description"`
	Review          string    `json:"review"`
	CreateAt        time.Time `json:"create_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	//DeletedAt       time.Time       `gorm:"index"  json:"-"`
	//Bookings        []Booking       `gorm:"foreignKey:PsikologID"`
	//BookingDetails  []BookingDetail `gorm:"foreignKey:PsikologID"`
}

type PsikologInput struct {
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Title           string `json:"title"`
	Price           int    `json:"price"`
	JenisKonsultasi string `json:"jenis_konsultasi"`
	Description     string `json:"description"`
	Review          string `json:"review"`
}

type UpdatePsikologInput struct {
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	Title           string `json:"title"`
	Price           int    `json:"price"`
	JenisKonsultasi string `json:"jenis_konsultasi"`
	Description     string `json:"description"`
}
