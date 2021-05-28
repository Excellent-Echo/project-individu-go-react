package account

import (
	"errors"
	"fmt"
	"project-individu-go-react/entities"
	"project-individu-go-react/helper"
	"time"
)

type Service interface {
	GetAllAccount() ([]AccountFormat, error)
	SaveNewAccount(account entities.AccountInput) (AccountFormat, error)
	GetAccountBySID(SID string) (AccountFormat, error)
	UpdateAccountBySID(SID string, dataInput entities.AccountUpdateInput) (AccountFormat, error)
	DeleteAccountBySID(SID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllAccount() ([]AccountFormat, error) {
	//Logic
	accounts, err := s.repository.FindAll()
	var accountFormat []AccountFormat

	for _, account := range accounts {
		formatAccount := FormatAccount(account)
		accountFormat = append(accountFormat, formatAccount)
	}

	if err != nil {
		return accountFormat, err
	}

	return accountFormat, nil
}

func (s *service) SaveNewAccount(account entities.AccountInput) (AccountFormat, error) {

	var newAccount = entities.Account{
		SID:         account.SID,
		Instrument:  account.Instrument,
		LastBalance: account.LastBalance,
		CID:         account.CID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Now(),
	}

	createAccount, err := s.repository.Create(newAccount)
	formatAccount := FormatAccount(createAccount)

	if err != nil {
		return formatAccount, err
	}

	return formatAccount, nil
}

func (s *service) GetAccountBySID(SID string) (AccountFormat, error) {
	if err := helper.ValidateIDNumber(SID); err != nil {
		return AccountFormat{}, err
	}

	account, err := s.repository.FindBySID(SID)

	if err != nil {
		return AccountFormat{}, err
	}

	if account.SID == 0 {
		newError := fmt.Sprintf("account sid %s not found", SID)
		return AccountFormat{}, errors.New(newError)
	}

	formatAccount := FormatAccount(account)

	return formatAccount, nil
}

func (s *service) UpdateAccountBySID(SID string, dataInput entities.AccountUpdateInput) (AccountFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(SID); err != nil {
		return AccountFormat{}, err
	}

	account, err := s.repository.FindBySID(SID)

	if err != nil {
		return AccountFormat{}, err
	}

	if account.SID == 0 {
		newError := fmt.Sprintf("account sid %s not found", SID)
		return AccountFormat{}, errors.New(newError)
	}

	if dataInput.Instrument != "" || len(dataInput.Instrument) != 0 {
		dataUpdate["instrument"] = dataInput.Instrument
	}

	if dataInput.LastBalance >= 0 {
		dataUpdate["last_balance"] = dataInput.LastBalance
	}

	dataUpdate["updated_at"] = time.Now()

	accountUpdate, err := s.repository.UpdateBySID(SID, dataUpdate)

	if err != nil {
		return AccountFormat{}, err
	}

	formatAccount := FormatAccount(accountUpdate)

	return formatAccount, nil
}

func (s *service) DeleteAccountBySID(SID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(SID); err != nil {
		return nil, err
	}

	account, err := s.repository.FindBySID(SID)

	if err != nil {
		return nil, err
	}

	if account.SID == 0 {
		newError := fmt.Sprintf("account sid %s not found", SID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteBySID(SID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete account SID: %s", SID)

	formatDelete := FormatDeleteAccount(msg)

	return formatDelete, nil
}
