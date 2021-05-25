package answer

import (
	"project-individu-go-react/entity"
	"time"
)

type AnswerService interface {
	PostNewAnswer(questionID int, userID int, answer entity.AnswerInput) (entity.Answers, error)
}

type answerService struct {
	repository AnswerRepository
}

func NewService(repository AnswerRepository) *answerService {
	return &answerService{repository}
}

func (s *answerService) PostNewAnswer(questionID int, userID int, answer entity.AnswerInput) (entity.Answers, error) {
	var newAnswer = entity.Answers{
		Content:    answer.Content,
		UserID:     uint32(userID),
		QuestionID: uint64(questionID),
		CreatedAt:  time.Now(),
	}

	createAnswer, err := s.repository.PostAnswer(newAnswer)

	if err != nil {
		return createAnswer, err
	}

	return createAnswer, nil
}
