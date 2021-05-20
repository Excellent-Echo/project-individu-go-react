import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreateAt  time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time `gorm:"index"`
}

type M021 struct {
	ID int `gorm:"primaryKey"`
}