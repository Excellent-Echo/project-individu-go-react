package entity

import (
	"time"

	"gorm.io/gorm"
)

type Tags struct {
	ID        uint32         `gorm:"primaryKey;autoIncrement;" json:"id"`
	Tag       string         `gorm:"size:255; not null;unique" json:"tag"`
	CreatedAt time.Time      `gorm:"type:datetime;not null;default:current_timestamp" json:"-"`
	Questions []Questions    `gorm:"many2many:question_tags;References:ID;constraint:OnUpdate:CASCADE" json:"questions"`
	Deleted   gorm.DeletedAt `json:"-"`
}
