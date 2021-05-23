package account

import "project-individu-go-react/entities"

type AccountFormat struct {
	SID         int    `json:"sid"`
	Instrument  string `json:"instrument"`
	LastBalance int    `json:"last_balance"`
	CID         int    `json:"cid"`
}

func FormatAccount(account entities.Account) AccountFormat {
	var formatAccount = AccountFormat{
		SID:         account.SID,
		Instrument:  account.Instrument,
		LastBalance: account.LastBalance,
		CID:         account.CID,
	}

	return formatAccount
}

type DeleteFormat struct {
	Message string `json:"message"`
}

func FormatDeleteAccount(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message: msg,
	}

	return deleteFormat
}
