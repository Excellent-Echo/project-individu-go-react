package entity

import "time"

type Post struct {
	PostID     uint64    `gorm:"primaryKey;autoIncrement;not null" json:"post_id"`
	Title      string    `gorm:"size:255;not null" json:"title"`
	Content    string    `gorm:"text;not null" json:"content"`
	UserID     uint32    `gorm:"not null" json:"user_id"`
	CategoryID uint32    `json:"category_id"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"deleted_at"`
	Comments   []Comment `gorm:"foreignKey:PostID" json:"comments"`
	Likes      []Like    `gorm:"foreignKey:PostID" json:"likes"`
}
