package entity

// UserInput untuk inputan yang diperlukan dari model users
type UserInput struct {
	RoleID    int    `json:"role" binding:"required"`
	Firstname string `json:"first_name" binding:"required"`
	Lastname  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
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
