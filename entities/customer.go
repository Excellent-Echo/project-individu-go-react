package entities

import "time"

type Costumer struct {
	CID       int       `gorm:"primaryKey" json:"cid"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Address   string    `json:"address"`
	Gender    string    `json:"gender"`
	NoPhone   string    `json:"no_phone"`
	UID       int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
	Accounts  []Account `gorm:"foreignKey:CID"`
}

type CostumerInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	NoPhone   string `json:"no_phone" binding:"required"`
	UID       int    `json:"id" binding:"required"`
}

type CostumerUpdateInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	NoPhone   string `json:"no_phone"`
}
