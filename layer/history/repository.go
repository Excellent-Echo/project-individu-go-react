package history

import (
	"project-individu-go-react/entities"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entities.History, error)
	Create(history entities.History) (entities.History, error)
	FindByHID(HID string) (entities.History, error)
	FindBySID(SID string) ([]entities.History, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//READ
func (r *repository) FindAll() ([]entities.History, error) {
	var histories []entities.History

	if err := r.db.Find(&histories).Error; err != nil {
		return histories, err
	}

	return histories, nil
}

// Read by hid
func (r *repository) FindByHID(HID string) (entities.History, error) {
	var history entities.History

	if err := r.db.Where("h_id = ?", HID).Find(&history).Error; err != nil {
		return history, err
	}

	return history, nil
}

//Read Account History
func (r *repository) FindBySID(SID string) ([]entities.History, error) {
	var histories []entities.History

	if err := r.db.Where("s_id = ?", SID).Find(&histories).Error; err != nil {
		return histories, err
	}

	return histories, nil
}

//CREATE
func (r *repository) Create(history entities.History) (entities.History, error) {
	if err := r.db.Create(&history).Error; err != nil {
		return history, err
	}

	return history, nil
}
