package customer

import (
	"errors"
	"fmt"
	"project-individu-go-react/entities"
	"project-individu-go-react/helper"
	"time"
)

type Service interface {
	GetAllCustomer() ([]CustomerFormat, error)
	SaveNewCustomer(customer entities.CostumerInput) (CustomerFormat, error)
	GetCustomerByCID(CID string) (CustomerFormat, error)
	UpdateCustomerByCID(CID string, dataInput entities.CostumerUpdateInput) (CustomerFormat, error)
	DeleteCustomerByCID(CID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllCustomer() ([]CustomerFormat, error) {
	//Logic
	customers, err := s.repository.FindAll()
	var customerFormat []CustomerFormat

	for _, customer := range customers {
		formatCustomer := FormatCustomer(customer)
		customerFormat = append(customerFormat, formatCustomer)
	}

	if err != nil {
		return customerFormat, err
	}

	return customerFormat, nil
}

func (s *service) SaveNewCustomer(customer entities.CostumerInput) (CustomerFormat, error) {

	var newCustomer = entities.Costumer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address:   customer.Address,
		Gender:    customer.Gender,
		NoPhone:   customer.NoPhone,
		UID:       customer.UID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	createCustomer, err := s.repository.Create(newCustomer)
	formatCustomer := FormatCustomer(createCustomer)

	if err != nil {
		return formatCustomer, err
	}

	return formatCustomer, nil
}

func (s *service) GetCustomerByCID(CID string) (CustomerFormat, error) {
	if err := helper.ValidateIDNumber(CID); err != nil {
		return CustomerFormat{}, err
	}

	customer, err := s.repository.FindByID(CID)

	if err != nil {
		return CustomerFormat{}, err
	}

	if customer.CID == 0 {
		newError := fmt.Sprintf("customer cid %s not found", CID)
		return CustomerFormat{}, errors.New(newError)
	}

	formatCustomer := FormatCustomer(customer)

	return formatCustomer, nil
}

func (s *service) UpdateCustomerByCID(CID string, dataInput entities.CostumerUpdateInput) (CustomerFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(CID); err != nil {
		return CustomerFormat{}, err
	}

	customer, err := s.repository.FindByID(CID)

	if err != nil {
		return CustomerFormat{}, err
	}

	if customer.CID == 0 {
		newError := fmt.Sprintf("user cid %s not found", CID)
		return CustomerFormat{}, errors.New(newError)
	}

	if dataInput.FirstName != "" || len(dataInput.FirstName) != 0 {
		dataUpdate["first_name"] = dataInput.FirstName
	}

	if dataInput.LastName != "" || len(dataInput.LastName) != 0 {
		dataUpdate["last_name"] = dataInput.LastName
	}

	if dataInput.Address != "" || len(dataInput.Address) != 0 {
		dataUpdate["address"] = dataInput.Address
	}

	if dataInput.Gender != "" || len(dataInput.Gender) != 0 {
		dataUpdate["gender"] = dataInput.Gender
	}

	if dataInput.NoPhone != "" || len(dataInput.NoPhone) != 0 {
		dataUpdate["no_phone"] = dataInput.NoPhone
	}

	dataUpdate["updated_at"] = time.Now()

	customerUpdate, err := s.repository.UpdateByID(CID, dataUpdate)

	if err != nil {
		return CustomerFormat{}, err
	}

	formatCustomer := FormatCustomer(customerUpdate)

	return formatCustomer, nil
}

func (s *service) DeleteCustomerByCID(CID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(CID); err != nil {
		return nil, err
	}

	customer, err := s.repository.FindByID(CID)

	if err != nil {
		return nil, err
	}

	if customer.CID == 0 {
		newError := fmt.Sprintf("user cid %s not found", CID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(CID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user CID: %S", CID)

	formatDelete := FormatDeleteCustomer(msg)

	return formatDelete, nil
}
