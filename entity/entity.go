package entity

import "time"

// User struct tabel users
type User struct {
	ID        int       `gorm:"primaryKey"`
	RoleID    int       `json:"role"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
	Booking   []Booking `gorm:"foreignKey:UserID"`
}

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
	ID             int `gorm:"primaryKey"`
	UserID         int
	PsikologID     int
	BookingDate    int
	BookingTime    int
	CreateAt       time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time       `gorm:"index"`
	BookingDetails []BookingDetail `gorm:"foreignKey:BookingID"`
}

// BookingDetail struct tabel bookingdetail
type BookingDetail struct {
	ID         int `gorm:"primaryKey"`
	BookingID  int
	PsikologID int
}

// Psikologi struct tabel psychologists
type Psikologi struct {
	ID              int `gorm:"primaryKey"`
	Firstname       string
	Lastname        string
	Title           string
	Price           int
	JenisKonsultasi string
	Description     string
	Review          string
	CreateAt        time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time       `gorm:"index"`
	Bookings        []Booking       `gorm:"foreignKey:PsikologID"`
	BookingDetails  []BookingDetail `gorm:"foreignKey:PsikologID"`
}
