package user

import "project-individu-go-react/entities"

type UserFormat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func FormatUser(user entities.User) UserFormat {
	var formatUser = UserFormat{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return formatUser
}

type DeleteFormat struct {
	Message string `json:"message"`
}

func FormatDeleteUser(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message: msg,
	}

	return deleteFormat
}
