package question

import (
	"project-individu-go-react/entity"
	"time"
)

type QuestionService interface {
	// FindAllQuestions() ([]entity.Questions, error)
	SaveNewQuestion(question entity.QuestionInput) (entity.Questions, error)
	// FindQuestionById(id string) (entity.Questions, error)
	// UpdateQuestionById(id string, dataInput entity.QuestionInput) (entity.Questions, error)
	// DeleteQuestionById(id string) (interface{}, error)
}

type questionService struct {
	repository QuestionRepository
}

func QuestionNewService(repository QuestionRepository) *questionService {
	return &questionService{repository}
}

func (s *questionService) SaveNewQuestion(question entity.QuestionInput) (entity.Questions, error) {
	var newQuestion = entity.Questions{
		Title:     question.Title,
		Content:   question.Content,
		TagID:     question.TagID,
		UserID:    question.UserID,
		CreatedAt: time.Now(),
	}

	createQuestion, err := s.repository.PostQuestion(newQuestion)

	if err != nil {
		return createQuestion, err
	}

	return createQuestion, nil
}
