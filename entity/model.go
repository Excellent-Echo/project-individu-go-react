package entity

import "time"

type User struct {
	UserID    uint32    `gorm:"primaryKey;autoIncrement;not null" json:"user_id"`
	FirstName string    `gorm:"size:255;not null" json:"first_name"`
	LastName  string    `gorm:"size:255;not null" json:"last_name"`
	UserName  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"updated_at"`
	Posts     []Post    `gorm:"foreignKey:UserID" json:"posts"`
	Comments  []Comment `gorm:"foreignKey:UserID" json:"comments"`
	Likes     []Like    `gorm:"foreignKey:UserID" json:"likes"`
}

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

type Category struct {
	CategoryID uint32    `gorm:"primaryKey;autoIncrement;not null" json:"category_id"`
	Category   string    `gorm:"size:255;not null" json:"category"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"updated_at"`
	Post       []Post    `gorm:"foreignKey:CategoryID" json:"post"`
}

type Comment struct {
	CommentID uint64    `gorm:"primaryKey;autoIncrement;not null" json:"comment_id"`
	Content   string    `gorm:"text;not null" json:"content"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	PostID    uint64    `gorm:"not null" json:"post_id"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"deleted_at"`
}

type Like struct {
	LikeID    uint64    `gorm:"primaryKey;autoIncrement" json:"like_id"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	PostID    uint64    `gorm:"not null" json:"post_id"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	DeletedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"deleted_at"`
}
