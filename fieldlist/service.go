package fieldlist

import (
	"errors"
	"fmt"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
	"strconv"
)

type Service interface {
	GetAllFieldList() ([]entity.FieldListSportName, error)
	SaveNewFieldList(fieldlist entity.FieldListInput, SportID string) (entity.FieldList, error)
	GetFieldListByID(fieldlistID string) (entity.FieldList, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllFieldList() ([]entity.FieldListSportName, error) {
	fieldlists, err := s.repository.FindAll()

	var formatFieldLists []entity.FieldListSportName

	for _, fieldlist := range fieldlists {
		formatFieldList := FormatFieldListSportName(fieldlist)
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

func (s *service) SaveNewFieldList(fieldlist entity.FieldListInput, SportID string) (entity.FieldList, error) {

	IDSport, _ := strconv.Atoi(SportID)

	var NewFieldList = entity.FieldList{
		FieldName:  fieldlist.FieldName,
		FieldImage: fieldlist.FieldImage,
		SportID:    IDSport,
	}

	createFieldList, err := s.repository.Create(NewFieldList)
	formatFieldList := FormatFieldList(createFieldList)

	if err != nil {
		return formatFieldList, err
	}

	return formatFieldList, nil
}
