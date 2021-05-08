package user

import (
	"project-individu-go-react/entity"
	"time"
)

// UserFormat untuk formatting yang ada di model user
type UserFormat struct {
	ID        int    `json:"id"`
	RoleID    int    `json:"role"`
	FIrstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatUser(user entity.User) UserFormat {
	var formatUser = UserFormat{
		ID:        user.ID,
		RoleID:    user.RoleID,
		FIrstname: user.Firstname,
		Lastname:  user.Lastname,
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
