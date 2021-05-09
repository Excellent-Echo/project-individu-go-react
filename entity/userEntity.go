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

// UserInput untuk inputan yang diperlukan dari model users
type UserInput struct {
	RoleID    int    `json:"role" binding:"required"`
	Firstname string `json:"first_name" binding:"required"`
	Lastname  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

// UserInputUpdate inputan yang diperlukan users untuk update data
type UserInputUpdate struct {
	RoleID    int    `json:"role"`
	FirstName string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
}
