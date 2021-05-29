package user

import (
	"project-individu-go-react/entity"
	"time"
)

type UserFormat struct {
	ID        int    `json:"id"`
	RoleID    int    `json:"role_id"`
	FIrstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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
		Password:  user.Password,
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
