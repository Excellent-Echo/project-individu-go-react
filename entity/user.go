package entity

import "time"

type User struct {
	User_id    int       `gorm:"primaryKey" json:"user_id"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"-"`
	Order_id   int       `gorm:"foreignKey:order_id"`
	// UserDetail []UserDetail `gorm:"foreignKey:UserID"`
}
