package bookinglist

import (
	"projectpenyewaanlapangan/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.BookingList, error)
	FindByID(ID string) (entity.BookingList, error)
	Create(bookinglist entity.BookingList) (entity.BookingList, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.BookingList, error)
	DeleteByID(ID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.BookingList, error) {
	var bookinglist []entity.BookingList

	if err := r.db.Find(&bookinglist).Error; err != nil {
		return bookinglist, err
	}

	return bookinglist, nil
}

func (r *repository) FindByID(ID string) (entity.BookingList, error) {
	var booking entity.BookingList

	if err := r.db.Where("id = ?", ID).Find(&booking).Error; err != nil {
		return booking, err
	}

	return booking, nil
}

func (r *repository) Create(bookinglist entity.BookingList) (entity.BookingList, error) {
	if err := r.db.Create(&bookinglist).Error; err != nil {
		return bookinglist, err
	}

	return bookinglist, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.BookingList, error) {

	var booking entity.BookingList

	if err := r.db.Model(&booking).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return booking, err
	}

	if err := r.db.Where("id = ?", ID).Find(&booking).Error; err != nil {
		return booking, err
	}

	return booking, nil
}

func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.BookingList{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
