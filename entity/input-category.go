package entity

type CategoryInput struct {
	CategoryName string `json:"category_name" binding:"required"`
}
