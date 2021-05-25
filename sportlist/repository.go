package sportlist

import (
	"projectpenyewaanlapangan/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.SportList, error)
	Create(sportlist entity.SportList) (entity.SportList, error)
	FindByID(ID string) (entity.SportList, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.SportList, error) {
	var sportlists []entity.SportList

	if err := r.db.Find(&sportlists).Error; err != nil {
		return sportlists, err
	}

	return sportlists, nil
}

func (r *repository) Create(sportlist entity.SportList) (entity.SportList, error) {
	if err := r.db.Create(&sportlist).Error; err != nil {
		return sportlist, err
	}

	return sportlist, nil
}

func (r *repository) FindByID(ID string) (entity.SportList, error) {
	var sportlist entity.SportList

	if err := r.db.Where("id = ?", ID).Find(&sportlist).Error; err != nil {
		return sportlist, err
	}

	return sportlist, nil
}
