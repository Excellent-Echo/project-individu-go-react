package entity

type KontakInput struct {
	NmKontak   string `json: "nmKontak"`
	NoTelp     string `json: "noTelp"`
	Perusahaan string `json: "perusahaan"`
	Email      string `json: "email"`
}
