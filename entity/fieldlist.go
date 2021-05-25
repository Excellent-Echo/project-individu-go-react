package entity

import "time"

type FieldList struct {
	ID         int         `gorm:"primaryKey" json:"id"`
	FieldName  string      `json:"field_name"`
	FieldImage string      `json:"field_image"`
	CreatedAt  time.Time   `json:"create_at"`
	UpdatedAt  time.Time   `json:"update_at"`
	DeletedAt  time.Time   `gorm:"index" json:"-"`
	SportID    []SportList `gorm:"foreignKey:FieldListID" json:"SportListID"`
}
