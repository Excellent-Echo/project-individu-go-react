package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint32         `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	FirstName string         `gorm:"size:255;not null" json:"first_name"`
	LastName  string         `gorm:"size:255;not null" json:"last_name"`
	UserName  string         `gorm:"size:255;not null;unique" json:"user_name"`
	Email     string         `gorm:"size:255;not null;unique" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time      `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:datetime;default:current_timestamp" json:"updated_at"`
	Questions []Questions    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"questions"`
	Answers   []Answers      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"answers"`
	Deleted   gorm.DeletedAt `json:"-"`
}
