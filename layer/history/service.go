package history

import (
	"errors"
	"fmt"
	"project-individu-go-react/entities"
	"project-individu-go-react/helper"
	"time"
)

type Service interface {
	GetAllHistory() ([]HistoryFormat, error)
	SaveNewHistory(history entities.InputHistory) (HistoryFormat, error)
	GetHistoriesBySID(SID string) ([]HistoryFormat, error)
	GetHistoryByHID(HID string) (HistoryFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllHistory() ([]HistoryFormat, error) {
	//Logic
	histories, err := s.repository.FindAll()
	var historyFormat []HistoryFormat

	for _, history := range histories {
		formatHistory := FormatHistory(history)
		historyFormat = append(historyFormat, formatHistory)
	}

	if err != nil {
		return historyFormat, err
	}

	return historyFormat, nil
}

func (s *service) SaveNewHistory(history entities.InputHistory) (HistoryFormat, error) {

	var newHistory = entities.History{
		HID:     history.HID,
		Credit:  history.Credit,
		Debit:   history.Debit,
		Balance: history.Balance,
		Date:    time.Now(),
		SID:     history.SID,
	}

	createHistory, err := s.repository.Create(newHistory)
	formatHistory := FormatHistory(createHistory)

	if err != nil {
		return formatHistory, err
	}

	return formatHistory, nil
}

func (s *service) GetHistoryByHID(HID string) (HistoryFormat, error) {
	if err := helper.ValidateIDNumber(HID); err != nil {
		return HistoryFormat{}, err
	}

	history, err := s.repository.FindByHID(HID)

	if err != nil {
		return HistoryFormat{}, err
	}

	if history.HID == 0 {
		newError := fmt.Sprintf("history hid %s not found", HID)
		return HistoryFormat{}, errors.New(newError)
	}

	formatHistory := FormatHistory(history)

	return formatHistory, nil
}

func (s *service) GetHistoriesBySID(SID string) ([]HistoryFormat, error) {
	var historyFormat []HistoryFormat
	if err := helper.ValidateIDNumber(SID); err != nil {
		return historyFormat, err
	}

	histories, err := s.repository.FindBySID(SID)

	if err != nil {
		return historyFormat, err
	}

	if len(histories) == 0 {
		newError := fmt.Sprintf("history from account sid %s not found", SID)
		return historyFormat, errors.New(newError)
	}

	for _, history := range histories {

		formatHistory := FormatHistory(history)
		historyFormat = append(historyFormat, formatHistory)
	}

	if err != nil {
		return historyFormat, err
	}

	return historyFormat, nil
}
