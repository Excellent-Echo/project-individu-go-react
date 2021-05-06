package entity

type UserInput struct {
	RoleID    int    `json:"role"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
