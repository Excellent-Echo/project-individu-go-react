package account

import (
	"project-individu-go-react/entities"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entities.Account, error)
	Create(account entities.Account) (entities.Account, error)
	UpdateBySID(SID string, dataUpdate map[string]interface{}) (entities.Account, error)
	DeleteBySID(SID string) (string, error)
	FindBySID(SID string) (entities.Account, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//READ
func (r *repository) FindAll() ([]entities.Account, error) {
	var accounts []entities.Account

	if err := r.db.Find(&accounts).Error; err != nil {
		return accounts, err
	}

	return accounts, nil
}

// Read by id
func (r *repository) FindBySID(SID string) (entities.Account, error) {
	var accounts entities.Account

	if err := r.db.Where("s_id = ?", SID).Find(&accounts).Error; err != nil {
		return accounts, err
	}

	return accounts, nil
}

//CREATE
func (r *repository) Create(account entities.Account) (entities.Account, error) {
	if err := r.db.Create(&account).Error; err != nil {
		return account, err
	}

	return account, nil
}

//UPDATE
func (r *repository) UpdateBySID(SID string, dataUpdate map[string]interface{}) (entities.Account, error) {
	var account entities.Account

	if err := r.db.Model(&account).Where("s_id = ?", SID).Updates(dataUpdate).Error; err != nil {
		return account, err
	}

	if err := r.db.Where("s_id = ?", SID).Find(&account).Error; err != nil {
		return account, err
	}

	return account, nil
}

//DELETE
func (r *repository) DeleteBySID(SID string) (string, error) {
	if err := r.db.Where("s_id = ?", SID).Delete(&entities.Account{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
