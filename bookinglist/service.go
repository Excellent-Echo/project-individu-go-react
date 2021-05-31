package bookinglist

import (
	"errors"
	"fmt"
	"projectpenyewaanlapangan/entity"
	"projectpenyewaanlapangan/helper"
	"time"
)

type Service interface {
	GetAllBookingList() ([]BookingFormat, error)
	SaveNewBooking(user entity.BookingList) (BookingFormat, error)
	GetBookingByID(bookingID string) (BookingFormat, error)
	UpdateBookByID(bookingID string, dataInput entity.BookingListInput) (BookingFormat, error)
	DeleteBookByID(bookingID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllBookingList() ([]BookingFormat, error) {
	//bisnis logic
	books, err := s.repository.FindAll()

	var formatBooks []BookingFormat

	for _, book := range books {
		formatBook := FormatBooking(book)
		formatBooks = append(formatBooks, formatBook)
	}

	if err != nil {
		return formatBooks, err
	}

	return formatBooks, nil
}

func (s *service) SaveNewBooking(book entity.BookingListInput) (BookingFormat, error) {
	var newBook = entity.BookingList{
		Date:      book.Date,
		User_id:   book.User_id,
		FirstName: book.FirstName,
		LastName:  book.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createBook, err := s.repository.Create(newBook)
	formatBook := FormatBooking(createBook)

	if err != nil {
		return formatBook, nil
	}

	return formatBook, nil
}

func (s *service) GetBookingByID(bookingID string) (BookingFormat, error) {
	book, err := s.repository.FindByID(bookingID)

	if err != nil {
		return BookingFormat{}, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("booking id %s not found", bookingID)
		return BookingFormat{}, errors.New(newError)
	}

	formatUser := FormatBooking(book)

	return formatUser, nil
}

func (s *service) UpdateBookByID(bookingID string, dataInput entity.BookingListInput) (BookingFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(bookingID); err != nil {
		return BookingFormat{}, err
	}

	book, err := s.repository.FindByID(bookingID)

	if err != nil {
		return BookingFormat{}, err
	}

	if book.ID == 0 {
		newError := fmt.Sprintf("booking id %s not found", bookingID)
		return BookingFormat{}, errors.New(newError)
	}

	// if dataInput.Date != time.Date("yyyy-mm-dd hh") || len(dataInput.Date.AppendFormat()) != 0 {
	// 	dataUpdate["date"] = dataInput.Date
	// }

	if dataInput.FirstName != "" || len(dataInput.FirstName) != 0 {
		dataUpdate["first_name"] = dataInput.FirstName
	}
	if dataInput.LastName != "" || len(dataInput.LastName) != 0 {
		dataUpdate["last_name"] = dataInput.LastName
	}

	dataUpdate["updated_at"] = time.Now()

	fmt.Println(dataUpdate)

	bookUpdated, err := s.repository.UpdateByID(bookingID, dataUpdate)

	if err != nil {
		return BookingFormat{}, err
	}

	formatBooking := FormatBooking(bookUpdated)

	return formatBooking, nil
}

func (s *service) DeleteBookByID(bookingID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(bookingID); err != nil {
		return nil, err
	}

	user, err := s.repository.FindByID(bookingID)

	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", bookingID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(bookingID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", bookingID)

	formatDelete := FormatDeleteUser(msg)

	return formatDelete, nil
}
