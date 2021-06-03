package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsAdmin   bool
	Events    []Event
}

type Event struct {
	gorm.Model
	EventName     string
	EventDate     time.Time
	EventLocation string
	Status        string
	UserID        int
	Participants  []User `gorm:"many2many:event_participants;"`
}
