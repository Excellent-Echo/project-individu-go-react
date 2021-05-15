package sportlist

import (
	"errors"
	"fmt"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
)

type Service interface {
	GetAllSportList() ([]entity.SportList, error)
	SaveNewSportList(sportlist entity.SportListInput) (entity.SportList, error)
	GetSportListByID(sportlistID string) (entity.SportList, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllSportList() ([]entity.SportList, error) {
	sportlists, err := s.repository.FindAll()

	var formatSportlists []entity.SportList

	for _, sportlist := range sportlists {
		formatSportlist := FormatSportlist(sportlist)
		formatSportlists = append(formatSportlists, formatSportlist)
	}

	if err != nil {
		return formatSportlists, err
	}

	return formatSportlists, nil
}

func (s *service) SaveNewSportList(sportlist entity.SportListInput) (entity.SportList, error) {

	var newSportList = entity.SportList{
		SportName: sportlist.SportName,
	}

	createSportList, err := s.repository.Create(newSportList)
	formatSportlist := FormatSportlist(createSportList)

	if err != nil {
		return formatSportlist, err
	}

	return formatSportlist, nil
}

func (s *service) GetSportListByID(sportlistID string) (entity.SportList, error) {
	if err := helper.ValidateIDNumber(sportlistID); err != nil {
		return entity.SportList{}, err
	}

	sportlist, err := s.repository.FindByID(sportlistID)

	if err != nil {
		return entity.SportList{}, err
	}

	if sportlist.ID == 0 {
		newError := fmt.Sprintf("sportlist id %s not found", sportlistID)
		return entity.SportList{}, errors.New(newError)
	}

	formatSportlist := FormatSportlist(sportlist)

	return formatSportlist, nil
}
