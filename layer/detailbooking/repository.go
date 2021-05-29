package detailbooking

import (
	"gorm.io/gorm"
	"project-individu-go-react/entity"
)

type Repository interface {
	FindAllBookingDetail() ([]entity.BookingDetail, error)
	CreateBookingDetail(user entity.BookingDetail) (entity.BookingDetail, error)
	FindBookingDetailByID(ID string) (entity.BookingDetail, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllBookingDetail() ([]entity.BookingDetail, error) {
	var bookingDetails []entity.BookingDetail
	if err := r.db.Find(&bookingDetails).Error; err != nil {
		return bookingDetails, err
	}
	return bookingDetails, nil
}

func (r *repository) CreateBookingDetail(bookingDet entity.BookingDetail) (entity.BookingDetail, error) {
	if err := r.db.Create(&bookingDet).Error; err != nil {
		return bookingDet, nil
	}
	return bookingDet, nil
}

func (r *repository) FindBookingDetailByID(ID string) (entity.BookingDetail, error) {
	var bookingDetailFind entity.BookingDetail
	if err := r.db.Where("id = ?", ID).Find(&bookingDetailFind).Error; err != nil {
		return bookingDetailFind, err
	}
	return bookingDetailFind, nil
}
