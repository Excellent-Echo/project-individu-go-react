package tag

import (
	"project-individu-go-react/entity"
)

type TagService interface {
	GetAllTags() ([]entity.Tags, error)
	SaveNewTag(tag entity.TagInput) (entity.Tags, error)
}

type tagService struct {
	repository TagRepository
}

func NewService(repository TagRepository) *tagService {
	return &tagService{repository}
}

func (s *tagService) GetAllTags() ([]entity.Tags, error) {
	tags, err := s.repository.GetAll()

	if err != nil {
		return tags, err
	}

	return tags, nil
}

func (s *tagService) SaveNewTag(tag entity.TagInput) (entity.Tags, error) {
	var newTag = entity.Tags{
		Tag: tag.Tag,
	}

	createTag, err := s.repository.NewTag(newTag)

	if err != nil {
		return createTag, err
	}

	return createTag, nil
}
