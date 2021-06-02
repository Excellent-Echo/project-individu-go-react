package migration

import "time"

type UserDetail struct {
	ID          int `gorm:"primaryKey"`
	NoHandphone uint
	Gender      string
	Address     string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
}

type User struct {
	ID          int           `gorm:"primaryKey" json:"id"`
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	Email       string        `json:"email"`
	Password    string        `json:"-"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   time.Time     `gorm:"index" json:"-"`
	UserDetail  []UserDetail  `gorm:"foreignKey:UserID"`
	UserProfile []UserProfile `gorm:"foreignKey:UserID"`
}

type UserProfile struct {
	ID          int    `json:"id"`
	ProfileUser string `json:"profile_user"` // kita tangkap dari file (foto) , path / dir file foto
	UserID      int    `json:"user_id"`
}

type BookingList struct {
	ID          int         `gorm:"primaryKey" json:"id"`
	Date        time.Time   `json:"date"`
	TimeForPlay int         `json:"time_for_play"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   time.Time   `gorm:"index" json:"-"`
	Users       []User      `gorm:"foreignKey:BookingListID"`
	Fields      []FieldList `gorm:"foreignKey:BookingListID"`
}

type FieldList struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	FieldName  string    `json:"field_name"`
	FieldImage string    `json:"field_image"`
	RentPrice  int       `json:"rent_price"`
	CreatedAt  time.Time `json:"create_at"`
	UpdatedAt  time.Time `json:"update_at"`
	DeletedAt  time.Time `gorm:"index" json:"-"`
}
