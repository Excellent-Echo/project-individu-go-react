package psikolog

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"time"
)

type Service interface {
	GetAllPsikolog() ([]PsikologFormat, error)
	SaveNewPsikolog(psikolog entity.PsikologInput) (PsikologFormat, error)
	GetPsikologByID(psikologID string) (PsikologFormat, error)
	DeletePsikologByID(psikologID string) (interface{}, error)
	UpdatePsikologByID(psikologID string, dataInput entity.UpdatePsikologInput) (PsikologFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllPsikolog() ([]PsikologFormat, error) {
	// melakukan business logic disini
	psikologs, err := s.repository.FindAll()
	var formatPsikologs []PsikologFormat

	for _, psikolog := range psikologs {
		formatPsikolog := FormatPsikolog(psikolog)
		formatPsikologs = append(formatPsikologs, formatPsikolog)
	}

	if err != nil {
		return formatPsikologs, err
	}

	return formatPsikologs, nil
}

func (s *service) SaveNewPsikolog(psikolog entity.PsikologInput) (PsikologFormat, error) {

	var newPsikolog = entity.Psikologi{
		Firstname:       psikolog.Firstname,
		Lastname:        psikolog.Lastname,
		Title:           psikolog.Title,
		Price:           psikolog.Price,
		JenisKonsultasi: psikolog.JenisKonsultasi,
		Description:     psikolog.Description,
		CreateAt:        time.Now(),
		UpdatedAt:       time.Now(),
	}

	createPsikolog, err := s.repository.Create(newPsikolog)
	formatPsikolog := FormatPsikolog(createPsikolog)

	if err != nil {
		return formatPsikolog, err
	}

	return formatPsikolog, nil
}

func (s *service) GetPsikologByID(psikologID string) (PsikologFormat, error) {
	if err := helper.ValidateIDNumber(psikologID); err != nil {
		return PsikologFormat{}, err
	}

	psikolog, err := s.repository.FindByID(psikologID)

	if err != nil {
		return PsikologFormat{}, err
	}

	if psikolog.ID == 0 {
		newError := fmt.Sprintf("psikolog id %s not found", psikologID)
		return PsikologFormat{}, errors.New(newError)
	}

	formatPsikolog := FormatPsikolog(psikolog)

	return formatPsikolog, nil
}

func (s *service) DeletePsikologByID(psikologID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(psikologID); err != nil {
		return nil, err
	}

	psikolog, err := s.repository.FindByID(psikologID)

	if err != nil {
		return nil, err
	}
	if psikolog.ID == 0 {
		newError := fmt.Sprintf("psikolog id %s not found", psikologID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(psikologID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete psikolog ID : %s", psikologID)

	formatDelete := FormatDeletePsikolog(msg)

	return formatDelete, nil
}

func (s *service) UpdatePsikologByID(psikologID string, dataInput entity.UpdatePsikologInput) (PsikologFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(psikologID); err != nil {
		return PsikologFormat{}, err
	}

	psikolog, err := s.repository.FindByID(psikologID)

	if err != nil {
		return PsikologFormat{}, err
	}

	if psikolog.ID == 0 {
		newError := fmt.Sprintf("psikolog id %s not found", psikologID)
		return PsikologFormat{}, errors.New(newError)
	}

	if dataInput.FirstName != "" || len(dataInput.FirstName) != 0 {
		dataUpdate["firstname"] = dataInput.FirstName
	}
	if dataInput.LastName != "" || len(dataInput.LastName) != 0 {
		dataUpdate["lastname"] = dataInput.LastName
	}
	if dataInput.Title != "" || len(dataInput.Title) != 0 {
		dataUpdate["title"] = dataInput.Title
	}
	if dataInput.Price != 0 {
		dataUpdate["price"] = dataInput.Price
	}
	if dataInput.JenisKonsultasi != "" || len(dataInput.JenisKonsultasi) != 0 {
		dataUpdate["jenis_konsultasi"] = dataInput.JenisKonsultasi
	}
	if dataInput.Description != "" || len(dataInput.Description) != 0 {
		dataUpdate["description"] = dataInput.Description
	}

	dataUpdate["updated_at"] = time.Now()

	fmt.Println(dataUpdate)

	psikologUpdated, err := s.repository.UpdateByID(psikologID, dataUpdate)

	if err != nil {
		return PsikologFormat{}, err
	}

	formatPsikolog := FormatPsikolog(psikologUpdated)

	return formatPsikolog, nil
}
