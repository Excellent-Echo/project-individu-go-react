package entity

import "time"

type Transaksi struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Times     time.Time `json:"times"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}
