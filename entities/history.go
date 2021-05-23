package entities

import "time"

type History struct {
	HID    int `gorm:"primaryKey" json:"hid"`
	Credit int `json:"credit"`
	Debit  int `json:"debit"`
	Date   time.Time
	SID    int `json:"sid"`
}
