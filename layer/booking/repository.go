package booking

import (
	"gorm.io/gorm"
	"project-individu-go-react/entity"
)

type Repository interface {
	FindAllBooking() ([]entity.Booking, error)
	CreateBooking(user entity.Booking) (entity.Booking, error)
	FindBookingByID(ID string) (entity.Booking, error)
	FindandDeleteBookingByID(ID string) (string, error)
	FindAndUpdateBookingByID(ID string, bookingUpdate map[string]interface{}) (entity.Booking, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllBooking() ([]entity.Booking, error) {
	var bookings []entity.Booking
	if err := r.db.Find(&bookings).Error; err != nil {
		return bookings, err
	}
	return bookings, nil
}

func (r *repository) CreateBooking(booking entity.Booking) (entity.Booking, error) {
	if err := r.db.Create(&booking).Error; err != nil {
		return booking, nil
	}
	return booking, nil
}

func (r *repository) FindBookingByID(ID string) (entity.Booking, error) {
	var bookingFind entity.Booking
	if err := r.db.Where("id = ?", ID).Find(&bookingFind).Error; err != nil {
		return bookingFind, err
	}
	return bookingFind, nil
}

func (r *repository) FindandDeleteBookingByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.Booking{}).Error; err != nil {
		return "error", err
	}
	return "success", nil
}

func (r *repository) FindAndUpdateBookingByID(ID string, bookingUpdate map[string]interface{}) (entity.Booking, error) {
	var bookingWillUpdate entity.Booking

	if err := r.db.Model(&bookingWillUpdate).Where("id = ?", ID).Updates(bookingUpdate).Error; err != nil {
		return bookingWillUpdate, err
	}

	if err := r.db.Where("id = ?", ID).Find(&bookingWillUpdate).Error; err != nil {
		return bookingWillUpdate, err
	}

	return bookingWillUpdate, nil
}
