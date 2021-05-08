package user

import "project-individu-go-react/entity"

type UserFormat struct {
	ID        int    `json:"id"`
	RoleID    int    `json:"role"`
	FIrstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
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
