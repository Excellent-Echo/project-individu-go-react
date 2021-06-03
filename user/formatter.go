package user

import "project-go-react/entity"

type UserFormat struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func FormattingUser(user entity.User) UserFormat {
	userFormat := UserFormat{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	return userFormat
}
