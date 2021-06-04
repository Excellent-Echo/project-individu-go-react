package entity

type Role struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
}

type Doctor struct {
	ID                  int `gorm:"primaryKey" json:"id"`
	UserID              int `gorm:"foreignKey:DoctorID" json:"user_id"`
	MedicalSpecialityID int `gorm:"foreignKey:DoctorID" json:"speciality_id"`
	AppointmentID       int `gorm:"foreignKey:DoctorID" json:"-"`
}
