package entity

import "time"

type Questions struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"text;not null" json:"content"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	TagID     uint32    `json:"tag_id"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:datetime" json:"deleted_at"`
	Answers   []Answers `gorm:"foreignKey:QuestionID" json:"answers"`
	Likes     []Likes   `gorm:"foreignKey:QuestionID" json:"likes"`
}
