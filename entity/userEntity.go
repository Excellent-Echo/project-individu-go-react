package entity

import "time"

// User struct untuk tabel users
type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	RoleID    int       `json:"role_id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
	Booking   []Booking `gorm:"foreignKey:UserID"`
}

// UserInput untuk inputan yang diperlukan dari model users
type UserInput struct {
	RoleID    int    `json:"role_id" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

// UserInputUpdate inputan yang diperlukan users untuk update data
type UserInputUpdate struct {
	RoleID    int    `json:"role_id"`
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
