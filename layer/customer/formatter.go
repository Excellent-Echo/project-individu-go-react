package customer

import "project-individu-go-react/entities"

type CustomerFormat struct {
	CID       int    `json:"cid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	NoPhone   string `json:"no_phone"`
	UID       int    `json:"id"`
}

func FormatCustomer(customer entities.Costumer) CustomerFormat {
	var formatCustomer = CustomerFormat{
		CID:       customer.CID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address:   customer.Address,
		Gender:    customer.Gender,
		NoPhone:   customer.NoPhone,
		UID:       customer.UID,
	}

	return formatCustomer
}

type DeleteFormat struct {
	Message string `json:"message"`
}

func FormatDeleteCustomer(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message: msg,
	}

	return deleteFormat
}
