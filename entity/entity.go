package entity

type Kontak struct {
	IdKontak   int    `gorm:"primaryKey" json: "idKontak"`
	NmKontak   string `json: "nmKontak"`
	NoTelp     string `json: "noTelp"`
	Perusahaan string `json: "perusahaan"`
	Email      string `json: "email"`
}
