package userdetail

import "project-individu-go-react/entity"

type UserDetailFormat struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Address     string `json:"address"`
	Gender      string `json:"gender"`
	NoHandphone int    `json:"no_handphone"`
	UserID      int    `json:"user_id"`
}

func FormatUserDetails(UserDetail entity.UserDetail) UserDetailFormat {
	var formatUserDetails = UserDetailFormat{
		ID:          UserDetail.ID,
		Address:     UserDetail.Address,
		Gender:      UserDetail.Gender,
		NoHandphone: UserDetail.NoHandphone,
		UserID:      UserDetail.UserID,
	}
	return formatUserDetails
}
