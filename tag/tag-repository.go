package tag

import (
	"project-individu-go-react/entity"

	"gorm.io/gorm"
)

type TagRepository interface {
	GetAll() ([]entity.Tags, error)
	NewTag(tag entity.Tags) (entity.Tags, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAll() ([]entity.Tags, error) {
	var Tags []entity.Tags

	err := r.db.Preload("Questions").Find(&Tags).Error
	if err != nil {
		return Tags, err
	}

	return Tags, nil
}

func (r *Repository) NewTag(tag entity.Tags) (entity.Tags, error) {
	if err := r.db.Create(&tag).Error; err != nil {
		return tag, err
	}

	return tag, nil
}
