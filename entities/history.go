package entities

import "time"

type History struct {
	HID     int `gorm:"primaryKey" json:"hid"`
	Credit  int `json:"credit"`
	Debit   int `json:"debit"`
	Balance int `json:"balance"`
	Date    time.Time
	SID     int `json:"sid"`
}

type InputHistory struct {
	HID     int `json:"hid" binding:"required"`
	Credit  int `json:"credit" binding:"required"`
	Debit   int `json:"debit" binding:"required"`
	Balance int `json:"balance" binding:"required"`
	SID     int `json:"sid" binding:"required"`
}
