package tabel

type Order struct {
	order_id   int `gorm:"primaryKey" json:"order_id"`
	product_id int `gorm:"foreignKey:product_id"`
}
