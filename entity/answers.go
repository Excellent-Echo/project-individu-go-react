package entity

import "time"

type Answers struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Content    string    `gorm:"text;not null" json:"content"`
	UserID     uint32    `gorm:"not null" json:"user_id"`
	QuestionID uint64    `gorm:"not null" json:"question_id"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"-"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"-"`
	// DeletedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"deleted_at"`
}
