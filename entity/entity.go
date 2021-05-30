package entity

type Product struct {
	ID          int    `gorm:"primaryKey" json:"id"`
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
