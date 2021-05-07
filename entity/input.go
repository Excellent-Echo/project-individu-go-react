package entity

type UserInput struct {
	RoleID    int    `json:"role"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type BookingInput struct {
	UserID      int `json:"user_id"`
	PsikologID  int `json:"psikolog_id"`
	BookingDate int `json:"booking_date"`
	BookingTime int `json:"booking_time"`
}

type PsikologInput struct {
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Title           string `json:"title"`
	Price           int    `json:"price"`
	JenisKonsultasi string `json:"jenis_konsultasi"`
	Description     string `json:"description"`
	Review          string `json:"review"`
}
