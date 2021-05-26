package entity

import "time"

type Category struct {
	category_id   int    `gorm:"primaryKey" json:"category_id"`
	nama_category string `json:"nama_category"`
}

type Order struct {
	order_id   int `gorm:"primaryKey" json:"order_id"`
	product_id int `gorm:"foreignKey:product_id"`
}

type Product struct {
	product_id         int    `gorm:"primaryKey" json:"product_id"`
	nama_product       string `json:"nama_product"`
	harga_product      int    `json:"harga_product"`
	dekskripsi_product string `json:"dekskripsi_product"`
	category_id        int    `gorm:"foreignKey:category_id"`
	gambar_product     string `json:"gambar_product"`
}

type User struct {
	user_id    int       `gorm:"primaryKey" json:"user_id"`
	first_name string    `json:"first_name"`
	last_name  string    `json:"last_name"`
	email      string    `json:"email"`
	password   string    `json:"-"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
	order_id   int       `gorm:"foreignKey:order_id"`
}
