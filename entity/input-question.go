package entity

type QuestionInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	// TagID   uint32 `json:"tag_id" binding:"required"`
	TagID  []Tags `json:"tags"`
	UserID uint32 `json:"user_id" binding:"required"`
}

type UpdateQuestionInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	TagID   uint32 `json:"tag_id"`
}
