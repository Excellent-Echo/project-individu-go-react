package user

import "project-individu-go-react/entity"

type UserFormat struct {
	User_id    int    `json:"user_id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}

func FormatUser(user entity.User) UserFormat {
	var formatUser = UserFormat{
		User_id:    user.User_id,
		First_name: user.First_name,
		Last_name:  user.Last_name,
		Email:      user.Email,
	}

	return formatUser
}
