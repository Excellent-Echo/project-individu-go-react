package entity

import "time"

// User struct untuk tabel users
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
