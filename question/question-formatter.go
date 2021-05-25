package question

import "project-individu-go-react/entity"

type QuestionFormat struct {
	ID      uint64           `json:"id"`
	Title   string           `json:"title"`
	Content string           `json:"content"`
	UserID  uint32           `json:"user_id"`
	Tags    []entity.Tags    `json:"tags"`
	Answers []entity.Answers `json:"answers"`
}

func FormattingQuestion(question entity.Questions) QuestionFormat {
	questionFormat := QuestionFormat{
		ID:      question.ID,
		Title:   question.Title,
		Content: question.Content,
		UserID:  question.UserID,
		Tags:    question.Tags,
		Answers: question.Answers,
	}

	return questionFormat
}
