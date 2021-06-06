package entity

type Product struct {
	Product_id         int    `gorm:"primaryKey" json:"product_id"`
	Nama_product       string `json:"nama_product"`
	Harga_product      int    `json:"harga_product"`
	Dekskripsi_product string `json:"dekskripsi_product"`
	Gambar_product     string `json:"gambar_product"`
	Category_id        int    `gorm:"foreignKey:category_id"`
}
