package tableschema

type Nota struct {
	IdNota          int    `gorm:"primaryKey json:"ID_Nota`
	IdLaundry       int    `gorm:"foreignKey:ID_Laundry`
	jenisPembayaran string `json:"Jenis_Pembayaran`
}
