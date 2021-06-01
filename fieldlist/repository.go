package fieldlist

import (
	"projectpenyewaanlapangan/entity"

	"gorm.io/gorm"
)

type Repository interface {
	// FindAll() ([]entity.FieldList, error)
	Create(fieldlist entity.FieldList) (entity.FieldList, error)
	FindByID(ID string) (entity.FieldList, error)
	FindAll() ([]entity.FieldList, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.FieldList, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(fieldlist entity.FieldList) (entity.FieldList, error) {
	if err := r.db.Create(&fieldlist).Error; err != nil {
		return fieldlist, err
	}

	return fieldlist, nil
}

func (r *repository) FindByID(ID string) (entity.FieldList, error) {
	var fieldlist entity.FieldList

	if err := r.db.Where("id = ?", ID).Find(&fieldlist).Error; err != nil {
		return fieldlist, err
	}

	return fieldlist, nil
}

func (e *repository) FindAll() ([]entity.FieldList, error) {
	var fieldLists []entity.FieldList

	if err := e.db.Find(&fieldLists).Error; err != nil {
		return fieldLists, err
	}

	return fieldLists, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.FieldList, error) {
	var fieldList entity.FieldList

	if err := r.db.Model(&fieldList).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return fieldList, err
	}

	if err := r.db.Where("id= ? ", ID).Find(&fieldList).Error; err != nil {
		return fieldList, err
	}

	return fieldList, nil
}
