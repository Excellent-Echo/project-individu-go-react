package question

import (
	"project-individu-go-react/entity"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	// FindAll() ([]entity.Questions, error)
	PostQuestion(question entity.Questions) (entity.Questions, error)
	// FindQuestion(id string) (entity.User, error)
	// UpdateQuestion(id string, dataUpdate map[string]interface{}) (entity.User, error)
	// DeleteQuestion(id string) (string, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) PostQuestion(question entity.Questions) (entity.Questions, error) {
	if err := r.db.Create(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}
