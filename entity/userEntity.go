package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	RoleID    int       `json:"role_id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//DeletedAt time.Time `gorm:"index" json:"-"`
	Booking []Booking `gorm:"foreignKey:UserID"`
}

type UserInput struct {
	RoleID    int    `json:"role_id" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type UserInputUpdate struct {
	RoleID    int    `json:"role_id"`
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
