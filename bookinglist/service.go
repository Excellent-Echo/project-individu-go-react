package bookinglist

import (
	"errors"
	"fmt"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/fieldlist"
	"projectpenyewaanlapangan/helper"
	"projectpenyewaanlapangan/user"
	"strconv"
	"time"
)

type Service interface {
	GetAllBookingList() ([]entity.BookingList, error)
	GetBookingByID(bookingID string) (entity.BookingList, error)
	GetAllBookingbyUserIDAndField(userID string, fieldID string) ([]entity.BookingList, error)
	SaveNewBooking(inputBooking entity.BookingListInput, userID string, fieldID string) (entity.BookingList, error)
	UpdateBookByID(bookingID string, dataInput entity.UpdateBookingListInput) (entity.BookingList, error)
}

type service struct {
	repository      Repository
	userRepository  user.Repository
	fieldRepository fieldlist.Repository
}

func NewService(repository Repository, userRepository user.Repository, fieldRepository fieldlist.Repository) *service {
	return &service{repository, userRepository, fieldRepository}
}

func (s *service) GetAllBookingList() ([]entity.BookingList, error) {
	books, err := s.repository.FindAll()

	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *service) GetAllBookingbyUserIDAndField(userID string, fieldID string) ([]entity.BookingList, error) {
	user, err := s.userRepository.FindByID(userID)
	field, err := s.fieldRepository.FindByID(fieldID)

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return []entity.BookingList{}, errors.New(newError)
	}

	if field.ID == 0 {
		newError := fmt.Sprintf("field id %s not found", fieldID)
		return []entity.BookingList{}, errors.New(newError)
	}

	books, err := s.repository.FindAllByUserAndField(userID, fieldID)

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

func (s *service) SaveNewBooking(inputBooking entity.BookingListInput, userID string, fieldID string) (entity.BookingList, error) {
	IDUser, _ := strconv.Atoi(userID)
	IDField, _ := strconv.Atoi(fieldID)

	var newBooking = entity.BookingList{
		Date:        inputBooking.Date,
		TimeForPlay: inputBooking.TimeForPlay,
		UserID:      IDUser,
		FieldID:     IDField,
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
