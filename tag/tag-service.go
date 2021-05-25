package tag

type TagService interface {
	GetAllTags() ([]TagFormat, error)
}

type tagService struct {
	repository TagRepository
}

func NewService(repository TagRepository) *tagService {
	return &tagService{repository}
}

func (s *tagService) GetAllTags() ([]TagFormat, error) {
	tags, err := s.repository.GetAll()

	var tagsFormat []TagFormat

	for _, tag := range tags {
		var tagFormat = FormattingTag(tag)
		tagsFormat = append(tagsFormat, tagFormat)
	}

	if err != nil {
		return tagsFormat, err
	}

	return tagsFormat, nil
}
