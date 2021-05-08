package entity

import "time"

// Role struct tabel role
type Role struct {
	ID       int    `gorm:"primaryKey"`
	Admin    string `json:"admin"`
	User     string `json:"user"`
	Psikolog string `json:"psikolog"`
	Users    []User `gorm:"foreignKey:RoleID"`
}

// Booking struct tabel bookings
type Booking struct {
	ID             int             `gorm:"primaryKey"`
	UserID         int             `json:"user_id"`
	PsikologID     int             `json:"psikolog_id"`
	BookingDate    int             `json:"booking_date"`
	BookingTime    int             `json:"booking_time"`
	CreateAt       time.Time       `json:"create_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      time.Time       `gorm:"index" json:"-"`
	BookingDetails []BookingDetail `gorm:"foreignKey:BookingID"`
}

// BookingDetail struct tabel bookingdetail
type BookingDetail struct {
	ID         int `gorm:"primaryKey"`
	BookingID  int `json:"booking_id"`
	PsikologID int `json:"psikolog_id"`
}

// Psikologi struct tabel psikolog
type Psikologi struct {
	ID              int             `gorm:"primaryKey"`
	Firstname       string          `json:"firstname"`
	Lastname        string          `json:"lastname"`
	Title           string          `json:"title"`
	Price           int             `json:"price"`
	JenisKonsultasi string          `json:"jenis_konsultasi"`
	Description     string          `json:"description"`
	Review          string          `json:"review"`
	CreateAt        time.Time       `json:"create_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       time.Time       `gorm:"index"  json:"-"`
	Bookings        []Booking       `gorm:"foreignKey:PsikologID"`
	BookingDetails  []BookingDetail `gorm:"foreignKey:PsikologID"`
}
