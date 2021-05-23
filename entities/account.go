package entities

import "time"

type Account struct {
	SID            int       `gorm:"primaryKey" json:"sid"`
	Instrument     string    `json:"instrument"`
	LastBalance    int       `json:"last_balance"`
	CID            int       `json:"cid"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:" deleted_at"`
	AccountHistory History   `gorm:"foreignKey:SID"`
}
