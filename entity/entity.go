package entity

import "time"

type Category struct {
	Category_id   int    `gorm:"primaryKey" json:"category_id"`
	Nama_category string `json:"nama_category"`
}

type Order struct {
	Order_id   int `gorm:"primaryKey" json:"order_id"`
	Product_id int `gorm:"foreignKey:product_id"`
}

type Product struct {
	Product_id         int    `gorm:"primaryKey" json:"product_id"`
	Nama_product       string `json:"nama_product"`
	Harga_product      int    `json:"harga_product"`
	Dekskripsi_product string `json:"dekskripsi_product"`
	Category_id        int    `gorm:"foreignKey:category_id"`
	Gambar_product     string `json:"gambar_product"`
}

type User struct {
	User_id    int       `gorm:"primaryKey" json:"user_id"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Order_id   int       `gorm:"foreignKey:order_id"`
}
