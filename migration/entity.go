package migration

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         int            `gorm:"primaryKey" json:"id"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	Email      string         `json:"email"`
	Password   string         `json:"-"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdateAt   time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Courses    []Course       `gorm:"foreignKey:UserID"`
	UserDetail UserDetail     `gorm:"foreignKey:UserID"`
}

type UserDetail struct {
	ID          int `gorm:"primaryKey"`
	Address     string
	Gender      string
	NoHandphone int
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Course struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	CategoryID  int
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	Courses     []Course `gorm:"foreignKey:CategoryID"`
}

type UserProfile struct {
	ID          int    `json:"id"`
	ProfileUser string `json:"profile_user"`
	UserID      int    `json:"user_id"`
}
