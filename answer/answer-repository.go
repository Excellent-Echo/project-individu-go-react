package answer

import (
	"project-individu-go-react/entity"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	PostAnswer(answer entity.Answers) (entity.Answers, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) PostAnswer(answer entity.Answers) (entity.Answers, error) {
	if err := r.db.Create(&answer).Error; err != nil {
		return answer, err
	}

	return answer, nil
}
