package entity

import "time"

type Comment struct {
	CommentID uint64    `gorm:"primaryKey;autoIncrement;not null" json:"comment_id"`
	Content   string    `gorm:"text;not null" json:"content"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	PostID    uint64    `gorm:"not null" json:"post_id"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"deleted_at"`
}
