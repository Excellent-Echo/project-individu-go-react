package entity

// Role struct tabel role
type Role struct {
	ID       int    `gorm:"primaryKey"`
	Admin    string `json:"admin"`
	User     string `json:"user"`
	Psikolog string `json:"psikolog"`
	Users    []User `gorm:"foreignKey:RoleID"`
}
