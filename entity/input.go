package entity

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type SportListInput struct {
	SportName string `json:"sport_name" binding:"required"`
}

type FieldListInput struct {
	FieldName  string `json:"field_name" binding:"required"`
	FieldImage string `json:"field_image" binding:"required"`
	SportID    string `json:"sport_id" binding:"required"`
}
