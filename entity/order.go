package entity

type Order struct {
	Order_id int       `gorm:"primaryKey" json:"order_id"`
	User_id  int       `json:"user_id"`
	Product  []Product `gorm:"foreignKey:product_id"`
}
