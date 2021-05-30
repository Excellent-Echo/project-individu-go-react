package migration

import "time"

type User struct {
	ID          int           `gorm:"primaryKey" json:"id"`
	RoleID      int           `json:"role_id"`
	Firstname   string        `json:"firstname"`
	Lastname    string        `json:"lastname"`
	Email       string        `json:"email"`
	Password    string        `json:"-"`
	CreateAt    time.Time     `json:"create_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   time.Time     `gorm:"index" json:"-"`
	Booking     []Booking     `gorm:"foreignKey:UserID"`
	UserProfile []UserProfile `gorm:"foreignKey:UserID"`
}

type Psikologi struct {
	ID              int             `gorm:"primaryKey" json:"id"`
	Firstname       string          `json:"firstname"`
	Lastname        string          `json:"lastname"`
	Title           string          `json:"title"`
	Price           int             `json:"price"`
	JenisKonsultasi string          `json:"jenis_konsultasi"`
	Description     string          `json:"description"`
	CreateAt        time.Time       `json:"create_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       time.Time       `gorm:"index"  json:"-"`
	Bookings        []Booking       `gorm:"foreignKey:PsikologID"`
	BookingDetails  []BookingDetail `gorm:"foreignKey:PsikologID"`
}

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

type BookingDetail struct {
	ID         int `gorm:"primaryKey"`
	BookingID  int `json:"booking_id"`
	PsikologID int `json:"psikolog_id"`
}
type Role struct {
	ID          int    `gorm:"primaryKey"`
	NamaRole    string `json:"nama_role"`
	Description string `json:"description"`
	Users       []User `gorm:"foreignKey:RoleID"`
}

type UserProfile struct {
	ID          int    `json:"id"`
	ProfileUser string `json:"profile_user"`
	UserID      int    `json:"user_id"`
}
