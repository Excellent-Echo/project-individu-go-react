package entity

type TagInput struct {
	Tag string `json:"tag" binding:"required"`
}
