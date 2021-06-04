package user

import (
	"doctorsAppointment/entity"
	"time"
)

type UserFormat struct {
	ID        int    `json:"id"`
	FIrstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatUser(user entity.User) UserFormat {
	var userFormat = UserFormat{
		ID:        user.ID,
		FIrstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return userFormat
}

func FormatDeleteUser(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
