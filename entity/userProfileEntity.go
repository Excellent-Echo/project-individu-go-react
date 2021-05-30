package entity

type UserProfile struct {
	ID        int    `json:"id"`
	ImageUser string `json:"image_user"`
	UserID    int    `json:"user_id"`
}
