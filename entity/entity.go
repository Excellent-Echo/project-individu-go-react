package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `gorm:"primaryKey" json:"id`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Product struct {
	ID          int    `gorm:"primaryKey" json:"id`
	ProductName string `json:"product_name"`
	Quantity    string `json:"quantity"`
	//User 		[]User `gorm:"foreignKey:ProductID"`
}

type Order struct {
	ID         int    `gorm:"primaryKey" json:"order_id"`
	OrderDate  uint   `json:"order_date"`
	ShipAdress string `json:"ship_adress"`
	//Product []Product `gorm:foreignKey:OrderID`
}
