package question

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"time"
)

type QuestionService interface {
	FindAllQuestions() ([]QuestionFormat, error)
	SaveNewQuestion(question entity.QuestionInput) (entity.Questions, error)
	FindQuestionById(id string) (QuestionFormat, error)
	UpdateQuestionById(id string, dataInput entity.UpdateQuestionInput) (entity.Questions, error)
	// DeleteQuestionById(id string) (interface{}, error)
}

type questionService struct {
	repository QuestionRepository
}

func QuestionNewService(repository QuestionRepository) *questionService {
	return &questionService{repository}
}

func (s *questionService) FindAllQuestions() ([]QuestionFormat, error) {
	questions, err := s.repository.FindAllQuestions()

	var questionsFormat []QuestionFormat

	for _, question := range questions {
		var questionFormat = FormattingQuestion(question)
		questionsFormat = append(questionsFormat, questionFormat)
	}

	if err != nil {
		return questionsFormat, err
	}

	return questionsFormat, nil
}

func (s *questionService) SaveNewQuestion(question entity.QuestionInput) (entity.Questions, error) {
	var newQuestion = entity.Questions{
		Title:   question.Title,
		Content: question.Content,
		Tags:    question.TagID,
		UserID:  question.UserID,
		// CreatedAt: time.Now(),
	}

	createQuestion, err := s.repository.PostQuestion(newQuestion)

	if err != nil {
		return createQuestion, err
	}

	return createQuestion, nil
}

func (s *questionService) FindQuestionById(id string) (QuestionFormat, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return QuestionFormat{}, err
	}

	question, err := s.repository.FindQuestionById(id)

	if err != nil {
		return QuestionFormat{}, err
	}

	if question.ID == 0 {
		newError := fmt.Sprintf("question id %s is not found", id)
		return QuestionFormat{}, errors.New(newError)
	}

	questionFormat := FormattingQuestion(question)

	return questionFormat, nil

}

func (s *questionService) UpdateQuestionById(id string, dataInput entity.UpdateQuestionInput) (entity.Questions, error) {
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
	// if len(dataInput.Tags) != 0 {
	// 	dataUpdate["tags"] = dataInput.Tags
	// }
	if dataInput.TagID != 0 {
		dataUpdate["tags"] = dataInput.TagID
	}

	dataUpdate["updated_at"] = time.Now()

	fmt.Println(dataUpdate)

	questionUpdate, err := s.repository.UpdateQuestion(id, dataUpdate)

	if err != nil {
		return entity.Questions{}, err
	}

	return questionUpdate, nil
}
