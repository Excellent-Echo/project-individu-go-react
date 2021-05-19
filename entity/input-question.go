package entity

type QuestionInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	TagID   uint32 `json:"tag_id" binding:"required"`
	UserID  uint32 `json:"user_id" binding:"required"`
}
