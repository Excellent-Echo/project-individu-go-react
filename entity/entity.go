package entity

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

type Product struct {
	ID          int    `gorm:"primaryKey"`
	ProductName string ``
	Quantity    string
	//User 		[]User `gorm:"foreignKey:ProductID"`
}

type Order struct {
	ID         int `gorm:"primaryKey"`
	OrderDate  uint
	ShipAdress string
	//Product []Product `gorm:foreignKey:OrderID`
}
