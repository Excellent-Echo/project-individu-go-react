package history

import (
	"project-individu-go-react/entities"
	"time"
)

type HistoryFormat struct {
	HID     int `json:"hid"`
	Credit  int `json:"credit"`
	Debit   int `json:"debit"`
	Balance int `json:"balance"`
	Date    time.Time
	SID     int `json:"sid"`
}

func FormatHistory(history entities.History) HistoryFormat {
	var formatHistory = HistoryFormat{
		HID:     history.HID,
		Credit:  history.Credit,
		Debit:   history.Debit,
		Balance: history.Balance,
		Date:    history.Date,
		SID:     history.SID,
	}

	return formatHistory
}
