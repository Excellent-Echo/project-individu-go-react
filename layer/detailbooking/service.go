package detailbooking

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
)

type Service interface {
	GetAllBookingDetail() ([]BookingDetailFormat, error)
	SaveNewBookingDetail(bookingDet entity.BookingDetailInput) (BookingDetailFormat, error)
	GetBookingDetailByID(bookingDetID string) (BookingDetailFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllBookingDetail() ([]BookingDetailFormat, error) {
	bookingDetails, err := s.repository.FindAllBookingDetail()
	var formatBookingDetails []BookingDetailFormat

	for _, bookingDetail := range bookingDetails {
		formatbookingdetail := FormatBookingDetails(bookingDetail)
		formatBookingDetails = append(formatBookingDetails, formatbookingdetail)
	}

	if err != nil {
		return formatBookingDetails, err
	}
	return formatBookingDetails, nil
}

func (s *service) SaveNewBookingDetail(bookingDet entity.BookingDetailInput) (BookingDetailFormat, error) {

	var newBookingDetail = entity.BookingDetail{
		BookingID:  bookingDet.BookingID,
		PsikologID: bookingDet.BookingID,
	}

	createBookingDetail, err := s.repository.CreateBookingDetail(newBookingDetail)
	formatBookingDetails := FormatBookingDetails(createBookingDetail)

	if err != nil {
		return formatBookingDetails, err
	}

	return formatBookingDetails, nil
}

func (s *service) GetBookingDetailByID(bookingDetID string) (BookingDetailFormat, error) {
	if err := helper.ValidateIDNumber(bookingDetID); err != nil {
		return BookingDetailFormat{}, err
	}

	bookingdetail, err := s.repository.FindBookingDetailByID(bookingDetID)

	if err != nil {
		return BookingDetailFormat{}, err
	}

	if bookingdetail.ID == 0 {
		newError := fmt.Sprintf("booking detail id %s not found", bookingDetID)
		return BookingDetailFormat{}, errors.New(newError)
	}

	formatBookingDetails := FormatBookingDetails(bookingdetail)

	return formatBookingDetails, nil
}
