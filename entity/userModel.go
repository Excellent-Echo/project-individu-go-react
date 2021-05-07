package entity

import "time"

type User struct {
	UserID    uint32    `gorm:"primaryKey;autoIncrement;not null" json:"user_id"`
	FirstName string    `gorm:"size:255;not null" json:"first_name"`
	LastName  string    `gorm:"size:255;not null" json:"last_name"`
	UserName  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;default:current_timestamp" json:"updated_at"`
	Posts     []Post    `gorm:"foreignKey:UserID" json:"posts"`
	Comments  []Comment `gorm:"foreignKey:UserID" json:"comments"`
	Likes     []Like    `gorm:"foreignKey:UserID" json:"likes"`
}
