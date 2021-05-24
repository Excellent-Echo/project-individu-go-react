package tableschema

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"ID_User`
	firstName string    `json:"first_name"`
	lastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	createdAt time.Time `json:"created_at`
	updatedAt time.Time `json:"updated_at`
}
