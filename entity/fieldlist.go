package entity

import "time"

type FieldList struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	FieldName  string    `json:"field_name"`
	FieldImage string    `json:"field_image"`
	RentPrice  int       `json:"rent_price"`
	CreatedAt  time.Time `json:"create_at"`
	UpdatedAt  time.Time `json:"update_at"`
	// DeletedAt  time.Time `gorm:"index" json:"-"`
	SportID int `json:"SportListID"`
}

type FieldListInput struct {
	FieldName  string `json:"field_name" binding:"required"`
	FieldImage string `json:"field_image" binding:"required"`
	RentPrice  int    `json:"rent_price" binding:"required"`
	SportID    int    `json:"sport_id" binding:"required"`
}

type FieldListSportName struct {
	ID         int       `json:"id"`
	FieldName  string    `json:"field_name"`
	FieldImage string    `json:"field_image"`
	RentPrice  int       `json:"rent_price" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"updated_at"`
	SportID    int       `json:"SportListID"`
	SportName  string    `json:"sport_name"`
}
