package entity

type Category struct {
	Category_id   int       `gorm:"primaryKey" json:"category_id"`
	Nama_category string    `json:"nama_category"`
	Product       []Product `gorm:"foreignKey:product_id"`
}
