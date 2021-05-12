package entity

import "time"

type Tags struct {
	TagID     uint32      `gorm:"primaryKey;autoIncrement;not null" json:"tag_id"`
	Tag       string      `gorm:"size:255;not null" json:"tag"`
	CreatedAt time.Time   `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time   `gorm:"type:datetime;not null;default:current_timestamp" json:"updated_at"`
	Questions []Questions `gorm:"foreignKey:TagID" json:"post"`
}
