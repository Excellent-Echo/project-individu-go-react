package entity

type SportList struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	SportName string `json:"sport_name"`
}

type SportListInput struct {
	SportName string `json:"sport_name" binding:"required"`
}
