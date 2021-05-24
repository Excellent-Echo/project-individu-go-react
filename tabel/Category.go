package tabel

type Category struct {
	category_id   int    `gorm:"primaryKey" json:"category_id"`
	nama_category string `json:"nama_category"`
}
