package entity

type Kontak struct {
	IdKontak   int    `gorm:"primaryKey" json: "idKontak"`
	NmKontak   string `json: "nmKontak"`
	NoTelp     string `json: "noTelp"`
	Perusahaan string `json: "perusahaan"`
	Email      string `json: "email"`
}

type Grup struct {
	IdGrup int    `gorm: "primaryKey" json:"idGrup"`
	NmGrup string `json: "nmGrup"`
}

type DtlGrup struct {
	IdDtlGrup int `gorm: "primaryKey" json:"idDtlGrup"`
	IdGrup    int `json:"idGrup"`
	IdKontak  int `json: "idKontak"`
}
