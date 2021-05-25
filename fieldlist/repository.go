package fieldlist

import (
	"projectpenyewaanlapangan/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.FieldList, error)
	Create(fieldlist entity.FieldList) (entity.FieldList, error)
	FindByID(ID string) (entity.FieldList, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.FieldList, error) {
	var fieldlists []entity.FieldList

	if err := r.db.Find(&fieldlists).Error; err != nil {
		return fieldlists, err
	}

	return fieldlists, nil
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
