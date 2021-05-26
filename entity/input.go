package tabel

import "time"

type User struct {
	user_id    int       `gorm:"primaryKey" json:"user_id"`
	first_name string    `json:"first_name"`
	last_name  string    `json:"last_name"`
	email      string    `json:"email"`
	password   string    `json:"-"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
	order_id   int       `gorm:"foreignKey:order_id"`
}
