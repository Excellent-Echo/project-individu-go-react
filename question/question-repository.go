package question

import (
	"project-individu-go-react/entity"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	FindAllQuestions() ([]entity.Questions, error)
	PostQuestion(question entity.Questions) (entity.Questions, error)
	FindQuestionById(id string) (entity.Questions, error)
	// UpdateQuestion(id string, dataUpdate map[string]interface{}) (entity.question, error)
	// DeleteQuestion(id string) (string, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) FindAllQuestions() ([]entity.Questions, error) {
	var Questions []entity.Questions

	err := r.db.Find(&Questions).Error
	if err != nil {
		return Questions, err
	}

	return Questions, nil
}

func (r *Repository) PostQuestion(question entity.Questions) (entity.Questions, error) {
	if err := r.db.Create(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}

func (r *Repository) FindQuestionById(id string) (entity.Questions, error) {
	var question entity.Questions

	if err := r.db.Where("id = ?", id).Find(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}
