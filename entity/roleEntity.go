package entity

type Role struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	NamaRole    string `json:"nama_role"`
	Description string `json:"description"`
	Users       []User `gorm:"foreignKey:RoleID"`
}

type RoleInput struct {
	NamaRole    string `json:"nama_role" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type RoleInputUpdate struct {
	NamaRole    string `json:"nama_role"`
	Description string `json:"description"`
}
