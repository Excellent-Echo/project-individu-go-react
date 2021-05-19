package psikolog

import (
	"gorm.io/gorm"
	"project-individu-go-react/entity"
)

type Repository interface {
	FindAll() ([]entity.Psikologi, error)
	Create(user entity.Psikologi) (entity.Psikologi, error)
	FindByID(ID string) (entity.Psikologi, error)
	DeleteByID(ID string) (string, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Psikologi, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Psikologi, error) {
	var psikologs []entity.Psikologi

	if err := r.db.Find(&psikologs).Error; err != nil {
		return psikologs, err
	}

	return psikologs, nil
}

func (r *repository) Create(psikolog entity.Psikologi) (entity.Psikologi, error) {
	if err := r.db.Create(&psikolog).Error; err != nil {
		return psikolog, err
	}

	return psikolog, nil
}

func (r *repository) FindByID(ID string) (entity.Psikologi, error) {
	var psikolog entity.Psikologi

	if err := r.db.Where("id = ?", ID).Find(&psikolog).Error; err != nil {
		return psikolog, err
	}

	return psikolog, nil
}

func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.Psikologi{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Psikologi, error) {

	var psikolog entity.Psikologi

	if err := r.db.Model(&psikolog).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return psikolog, err
	}

	if err := r.db.Where("id = ?", ID).Find(&psikolog).Error; err != nil {
		return psikolog, err
	}

	return psikolog, nil
}
