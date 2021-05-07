package entity

import "time"

type Category struct {
	CategoryID uint32    `gorm:"primaryKey;autoIncrement;not null" json:"category_id"`
	Category   string    `gorm:"size:255;not null" json:"category"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"updated_at"`
	Post       []Post    `gorm:"foreignKey:CategoryID" json:"post"`
}
