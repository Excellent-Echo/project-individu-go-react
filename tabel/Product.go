package tabel

type Product struct {
	product_id         int    `gorm:"primaryKey" json:"product_id"`
	nama_product       string `json:"nama_product"`
	harga_product      int    `json:"harga_product"`
	dekskripsi_product string `json:"dekskripsi_product"`
	category_id        int    `gorm:"foreignKey:category_id"`
	gambar_product     string `json:"gambar_product"`
}
