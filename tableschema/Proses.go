package tableschema

type Proses struct {
	IdProses  int `gorm:"primaryKey json:"ID_Proses`
	IdLaundry int `gorm:"foreignKey:ID_Laundry`
}
