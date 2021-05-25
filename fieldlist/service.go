package fieldlist

import (
	"errors"
	"fmt"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
)

type Service interface {
	GetAllFieldList() ([]entity.FieldList, error)
	SaveNewFieldList(fieldlist entity.FieldList) (entity.FieldList, error)
	GetFieldListByID(fieldlistID string) (entity.FieldList, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllFieldList() ([]entity.FieldList, error) {
	fieldlists, err := s.repository.FindAll()

	var formatFieldLists []entity.FieldList

	for _, fieldlist := range fieldlists {
		formatFieldList := FormatFieldList(fieldlist)
		formatFieldLists = append(formatFieldLists, formatFieldList)
	}

	if err != nil {
		return formatFieldLists, err
	}

	return formatFieldLists, nil
}

func (s *service) GetFieldListByID(fieldlistID string) (entity.FieldList, error) {
	if err := helper.ValidateIDNumber(fieldlistID); err != nil {
		return entity.FieldList{}, err
	}

	fieldlist, err := s.repository.FindByID(fieldlistID)

	if err != nil {
		return entity.FieldList{}, err
	}

	if fieldlist.ID == 0 {
		newError := fmt.Sprintf("field list id %s not found", fieldlistID)
		return entity.FieldList{}, errors.New(newError)
	}

	formatFieldList := FormatFieldList(fieldlist)

	return formatFieldList, nil
}
