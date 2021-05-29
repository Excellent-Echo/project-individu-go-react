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
	SaveNewFieldList(pathFile string, SportID int) (entity.FieldListSportName, error)
	GetFieldListByID(fieldlistID string) (entity.FieldListSportName, error)
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

func (s *service) GetFieldListByID(fieldlistID string) (entity.FieldListSportName, error) {
	if err := helper.ValidateIDNumber(fieldlistID); err != nil {
		return entity.FieldListSportName{}, err
	}

	fieldlist, err := s.repository.FindByID(fieldlistID)

	if err != nil {
		return entity.FieldListSportName{}, err
	}

	if fieldlist.ID == 0 {
		newError := fmt.Sprintf("field list id %s not found", fieldlistID)
		return entity.FieldListSportName{}, errors.New(newError)
	}

	formatFieldList := FormatFieldListSportName(fieldlist)

	return formatFieldList, nil
}

// func (s *service) SaveNewFieldList(fieldlist entity.FieldListInput, SportID string) (entity.FieldList, error) {

// 	IDSport, _ := strconv.Atoi(SportID)

// 	var NewFieldList = entity.FieldList{
// 		FieldName:  fieldlist.FieldName,
// 		FieldImage: fieldlist.FieldImage,
// 		RentPrice:  fieldlist.RentPrice,
// 		SportID:    IDSport,
// 	}

// 	createFieldList, err := s.repository.Create(NewFieldList)
// 	formatFieldList := FormatFieldList(createFieldList)

// 	if err != nil {
// 		return formatFieldList, err
// 	}

// 	return formatFieldList, nil
// }

func (s *service) SaveNewFieldList(pathFile string, SportID int) (entity.FieldListSportName, error) {
	ID := strconv.Itoa(SportID)

	fieldImageCheck, _ := s.repository.FindByID(ID)

	if fieldImageCheck.SportID == SportID {
		errorStatus := fmt.Sprintf("field list for sport id %d has been created", SportID)
		return fieldImageCheck, errors.New(errorStatus)
	}

	newFieldImage := entity.FieldListSportName{
		FieldName:  fieldImageCheck.FieldName,
		FieldImage: fieldImageCheck.FieldImage,
		RentPrice:  fieldImageCheck.RentPrice,
		SportID:    fieldImageCheck.SportID,
	}

	fieldImage, err := s.repository.Create(newFieldImage)

	if err != nil {
		return fieldImage, err
	}

	return fieldImage, err

}
