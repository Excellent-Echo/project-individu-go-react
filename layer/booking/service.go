package booking

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"time"
)

type Service interface {
	GetAllBooking() ([]BookingFormat, error)
	SaveNewBooking(booking entity.BookingInput) (BookingFormat, error)
	GetBookingByID(bookingID string) (BookingFormat, error)
	GetandDeleteBookingByID(bookingID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) GetAllBooking() ([]BookingFormat, error) {
	bookings, err := s.repository.FindAllBooking()
	var formatBookings []BookingFormat

	for _, booking := range bookings {
		formatbooking := FormatBooking(booking)
		formatBookings = append(formatBookings, formatbooking)
	}

	if err != nil {
		return formatBookings, err
	}
	return formatBookings, nil
}

func (s *service) SaveNewBooking(booking entity.BookingInput) (BookingFormat, error) {
	var newBooking = entity.Booking{
		UserID:      booking.UserID,
		PsikologID:  booking.PsikologID,
		BookingDate: booking.BookingDate,
		BookingTime: booking.BookingTime,
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}

	createBooking, err := s.repository.CreateBooking(newBooking)
	formatBooking := FormatBooking(createBooking)

	if err != nil {
		return formatBooking, err
	}

	return formatBooking, nil
}

func (s *service) GetBookingByID(bookingID string) (BookingFormat, error) {
	if err := helper.ValidateIDNumber(bookingID); err != nil {
		return BookingFormat{}, err
	}

	booking, err := s.repository.FindBookingByID(bookingID)

	if err != nil {
		return BookingFormat{}, err
	}

	if booking.ID == 0 {
		newError := fmt.Sprintf("booking id %s not found", bookingID)
		return BookingFormat{}, errors.New(newError)
	}

	formatBooking := FormatBooking(booking)

	return formatBooking, nil
}

func (s *service) GetandDeleteBookingByID(bookingID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(bookingID); err != nil {
		return BookingFormat{}, err
	}

	booking, err := s.repository.FindBookingByID(bookingID)

	if err != nil {
		return nil, err
	}

	if booking.ID == 0 {
		newError := fmt.Sprintf("booking id %s not found", bookingID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.FindandDeleteBookingByID(bookingID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	message := fmt.Sprintf("success delete booking by ID : %s", bookingID)

	formatDelete := FormatDeleteBooking(message)

	return formatDelete, nil
}
