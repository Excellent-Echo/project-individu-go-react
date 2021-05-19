package entity

import "time"

type Likes struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint32    `gorm:"not null" json:"user_id"`
	QuestionID uint64    `gorm:"not null" json:"question_id"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	DeletedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"deleted_at"`
}
