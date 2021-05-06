package entity

type UserInput struct {
	RoleID    int    `json:"role"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type BookingInput struct {
	BookingDate int `json:"booking_date"`
	BookingTime int `json:"booking_time"`
}
