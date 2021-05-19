package entity

import "time"

type Tags struct {
	ID        uint32      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Tag       string      `gorm:"size:255;not null" json:"tag"`
	CreatedAt time.Time   `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	Questions []Questions `gorm:"foreignKey:TagID" json:"post"`
}
