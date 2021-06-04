package entity

// type Role struct {
// 	ID          int `gorm:"primaryKey"`
// 	Title       string
// 	Description string
// }

// type MedicalSpecialist struct {
// 	ID          int `gorm:"primaryKey"`
// 	Title       string
// 	Description string
// }

// type Appointment struct {
// 	ID          int `gorm:"primaryKey"`
// 	Time        time.Time
// 	Date        time.Time
// 	IsCompleted bool
// }

// type Doctor struct {
// 	ID                int                 `gorm:"primaryKey"`
// 	User              []User              `gorm:"foreignKey:DoctorID"`
// 	MedicalSpecialist []MedicalSpecialist `gorm:"foreignKey:DoctorID"`
// 	Appointment       []Appointment       `gorm:"foreignKey:DoctorID"`
// }

// type Patient struct {
// 	ID          int           `gorm:"primaryKey"`
// 	User        []User        `gorm:"foreignKey:PatientID"`
// 	Appointment []Appointment `gorm:"foreignKey:PatientID"`
// }
