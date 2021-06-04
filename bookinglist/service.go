package bookinglist

import (
	"errors"
	"fmt"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
	"time"
)

type Service interface {
	GetAllBookingList() ([]entity.BookingList, error)
	GetBookingByID(bookingID string) (entity.BookingList, error)
	SaveNewBooking(inputBooking entity.BookingListInput) (entity.BookingList, error)
	UpdateBookByID(bookingID string, dataInput entity.UpdateBookingListInput) (entity.BookingList, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllBookingList() ([]entity.BookingList, error) {
	books, err := s.repository.FindAll()

	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *service) GetBookingByID(bookingID string) (entity.BookingList, error) {
	book, err := s.repository.FindByID(bookingID)

	if book.ID == 0 {
		newError := fmt.Sprintf("booking id %s not found", bookingID)
		return entity.BookingList{}, errors.New(newError)
	}

	if err != nil {
		return book, nil
	}

	return book, nil
}

func (s *service) SaveNewBooking(inputBooking entity.BookingListInput) (entity.BookingList, error) {

	var newBooking = entity.BookingList{
		Date:        inputBooking.Date,
		TimeForPlay: inputBooking.TimeForPlay,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createBook, err := s.repository.Create(newBooking)

	if err != nil {
		return createBook, err
	}

	return createBook, nil
}

func (s *service) UpdateBookByID(bookingID string, dataInput entity.UpdateBookingListInput) (entity.BookingList, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(bookingID); err != nil {
		return entity.BookingList{}, err
	}

	book, err := s.repository.FindByID(bookingID)

	if err != nil {
		return entity.BookingList{}, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("booking id %s not found", bookingID)
		return entity.BookingList{}, errors.New(newError)
	}

	if dataInput.TimeForPlay != 0 {
		dataUpdate["time_for_play"] = dataInput.TimeForPlay
	}

	dataUpdate["date"] = dataInput.Date

	bookUpdated, err := s.repository.UpdateByID(bookingID, dataUpdate)

	if err != nil {
		return bookUpdated, err
	}

	return bookUpdated, nil
}
