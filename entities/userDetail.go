package entities

import "time"

type UserDetail struct {
	ID        int `gorm:"primaryKey"`
	NoPhone   uint
	Gender    string
	Address   string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"index"`
}
