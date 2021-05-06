package entity

import "time"

type User struct {
	UserID    int       `gorm:"primaryKey;autoIncrement;not null" json:"user_id"`
	FirstName string    `gorm:"size:255;not null" json:"first_name"`
	LastName  string    `gorm:"size:255;not null" json:"last_name"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Post      []Post    `gorm:"foreignKey:UserID" json:"posts"`
	Comments  []Comment `gorm:"foreignKey:PostID" json:"comments"`
	Likes     []Like    `gorm:"foreignKey:PostID" json:"likes"`
}

type Post struct {
	PostID     int       `gorm:"primaryKey;autoIncrement;not null" json:"post_id"`
	Title      string    `gorm:"size:255;not null" json:"title"`
	Content    string    `gorm:"text;not null" json:"content"`
	UserID     int       `gorm:"not null" json:"user_id"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
	Comments   []Comment `gorm:"foreignKey:PostID" json:"comments"`
	Likes      []Like    `gorm:"foreignKey:PostID" json:"likes"`
}

type Category struct {
	CategoryID int       `gorm:"primaryKey;autoIncrement;not null" json:"category_id"`
	Category   string    `gorm:"size:255;not null" json:"category"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Post       []Post    `gorm:"foreignKey:UserID" json:"post"`
}

type Comment struct {
	CommentID int       `gorm:"primaryKey;autoIncrement;not null" json:"comment_id"`
	Content   string    `gorm:"text;not null" json:"content"`
	UserID    int       `gorm:"not null" json:"user_id"`
	PostID    int       `gorm:"not null" json:"post_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

type Like struct {
	LikeID    int       `gorm:"primaryKey;autoIncrement" json:"like_id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	PostID    int       `gorm:"not null" json:"post_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	DeletedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
}
