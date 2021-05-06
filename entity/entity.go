package entity

import "time"

type User struct {
	IdUSer    int       `gorm:"primaryKey" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

type SportList struct {
	IdSport   int    `gorm:"primaryKey" json:"id"`
	SportName string `json:"sport_name"`
}
