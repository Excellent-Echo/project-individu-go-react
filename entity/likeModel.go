package entity

import "time"

type Like struct {
	LikeID    uint64    `gorm:"primaryKey;autoIncrement" json:"like_id"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	PostID    uint64    `gorm:"not null" json:"post_id"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	DeletedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"deleted_at"`
}
