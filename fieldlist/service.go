package fieldlist

import (
	"errors"
	"fmt"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
	"time"
)

type Service interface {
	GetAllFieldList() ([]entity.FieldList, error)
	SaveNewFieldList(pathFile string, fieldListInput entity.FieldListInput) (FieldListFormat, error)
	GetFieldListByID(fieldlistID string) (entity.FieldList, error)
	UpdateFieldListById(pathFile string, fieldlistID string, fieldListInput entity.FieldListInput) (FieldListFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllFieldList() ([]entity.FieldList, error) {
	fieldList, err := s.repository.FindAll()

	if err != nil {
		return fieldList, err
	}

	return fieldList, nil
}

func (s *service) GetFieldListByID(fieldlistID string) (entity.FieldList, error) {
	fieldList, err := s.repository.FindByID(fieldlistID)

	if err != nil {
		return fieldList, err
	}

	return fieldList, nil
}

func (s *service) SaveNewFieldList(pathFile string, fieldListInput entity.FieldListInput) (FieldListFormat, error) {

	newFieldList := entity.FieldList{
		FieldName:  fieldListInput.FieldName,
		RentPrice:  fieldListInput.RentPrice,
		FieldImage: pathFile,
	}

	createFieldList, err := s.repository.Create(newFieldList)
	formatFieldList := FormatFieldList(createFieldList)

	if err != nil {
		return formatFieldList, err
	}

	return formatFieldList, nil
}

func (s *service) UpdateFieldListById(pathFile string, fieldlistID string, fieldListInput entity.FieldListInput) (FieldListFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(fieldlistID); err != nil {
		return FieldListFormat{}, err
	}

	field, err := s.repository.FindByID(fieldlistID)

	if err != nil {
		return FieldListFormat{}, err
	}

	if field.ID == 0 {
		newError := fmt.Sprintf("field id %s not found", fieldlistID)
		return FieldListFormat{}, errors.New(newError)
	}

	if fieldListInput.FieldName != "" || len(fieldListInput.FieldName) != 0 {
		dataUpdate["field_name"] = fieldListInput.FieldName
	}

	if fieldListInput.RentPrice != 0 {
		dataUpdate["rent_price"] = fieldListInput.RentPrice
	}

	dataUpdate["field_image"] = pathFile

	dataUpdate["updated_at"] = time.Now()

	fieldUpdated, err := s.repository.UpdateByID(fieldlistID, dataUpdate)

	if err != nil {
		return FieldListFormat{}, err
	}

	formatFieldList := FormatFieldList(fieldUpdated)

	return formatFieldList, nil
}
