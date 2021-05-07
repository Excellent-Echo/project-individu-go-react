package user

import "project-individu-go-react/entity"

type UserFormat struct {
	UserID    uint32 `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func FormattingUser(user entity.User) UserFormat {
	userFormat := UserFormat{
		UserID:    user.UserID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return userFormat
}
