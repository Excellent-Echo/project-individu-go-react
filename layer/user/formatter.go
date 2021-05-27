package user

import (
	"project-individu-go-react/entity"
	"time"
)

type UserFormat struct {
	UserID    uint32 `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"delete_time"`
}

func FormattingUser(user entity.User) UserFormat {
	userFormat := UserFormat{
		UserID:    user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return userFormat
}

func FormatDelete(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}
	return deleteFormat
}
