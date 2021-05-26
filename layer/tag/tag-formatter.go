package tag

import "project-individu-go-react/entity"

type TagFormat struct {
	ID  uint32 `json:"id"`
	Tag string `json:"tag"`
}

func FormattingTag(tag entity.Tags) TagFormat {
	tagFormat := TagFormat{
		ID:  tag.ID,
		Tag: tag.Tag,
	}

	return tagFormat
}
