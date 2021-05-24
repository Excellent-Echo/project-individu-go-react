package question

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"time"
)

type QuestionService interface {
	FindAllQuestions() ([]entity.Questions, error)
	SaveNewQuestion(question entity.QuestionInput) (entity.Questions, error)
	FindQuestionById(id string) (entity.Questions, error)
	UpdateQuestionById(id string, dataInput entity.QuestionInput) (entity.Questions, error)
	// DeleteQuestionById(id string) (interface{}, error)
}

type questionService struct {
	repository QuestionRepository
}

func QuestionNewService(repository QuestionRepository) *questionService {
	return &questionService{repository}
}

func (s *questionService) FindAllQuestions() ([]entity.Questions, error) {
	questions, err := s.repository.FindAllQuestions()

	if err != nil {
		return questions, err
	}

	return questions, nil
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

func (s *questionService) FindQuestionById(id string) (entity.Questions, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.Questions{}, err
	}

	question, err := s.repository.FindQuestionById(id)

	if err != nil {
		return entity.Questions{}, err
	}

	if question.ID == 0 {
		newError := fmt.Sprintf("question id %s is not found", id)
		return entity.Questions{}, errors.New(newError)
	}

	return question, nil

}

func (s *questionService) UpdateQuestionById(id string, dataInput entity.QuestionInput) (entity.Questions, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.Questions{}, err
	}

	question, err := s.repository.FindQuestionById(id)

	if err != nil {
		return entity.Questions{}, err
	}

	if question.ID == 0 {
		newError := fmt.Sprintf("question id %s is not found", id)
		return entity.Questions{}, errors.New(newError)
	}

	if dataInput.Title != "" || len(dataInput.Title) != 0 {
		dataUpdate["title"] = dataInput.Title
	}
	if dataInput.Content != "" || len(dataInput.Content) != 0 {
		dataUpdate["content"] = dataInput.Content
	}
	if dataInput.TagID != 0 {
		dataUpdate["tag_id"] = dataInput.TagID
	}

	dataUpdate["updated_at"] = time.Now()

	fmt.Println(dataUpdate)

	questionUpdate, err := s.repository.UpdateQuestion(id, dataUpdate)

	if err != nil {
		return entity.Questions{}, err
	}

	return questionUpdate, nil
}
