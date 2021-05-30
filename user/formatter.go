package user

import (
	"project-go-react/entity"
	"time"
)

type UserFormat struct {
	ID        int    `json: "id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json: "email"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatUser(user entity.User) UserFormat {
	var formatUser = UserFormat{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	return formatUser

}

func FormatDeleteUser(message string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    message,
		TimeDelete: time.Now(),
	}
	return deleteFormat
}
