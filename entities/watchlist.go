package entities

type Watchlist struct {
	ID         int `gorm:"primaryKey"`
	EmitenCode string
	UserID     int
}
