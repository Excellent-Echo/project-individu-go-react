package booking

import (
	"project-individu-go-react/entity"
	"time"
)

type BookingFormat struct {
	ID          int `json:"id"`
	PsikologID  int `json:"psikolog_id"`
	BookingDate int `json:"booking_date"`
	BookingTime int `json:"booking_time"`
}

func FormatBooking(booking entity.Booking) BookingFormat {
	var formatBooking = BookingFormat{
		ID:          booking.ID,
		PsikologID:  booking.PsikologID,
		BookingDate: booking.BookingDate,
		BookingTime: booking.BookingTime,
	}
	return formatBooking
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatDeleteBooking(message string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    message,
		TimeDelete: time.Now(),
	}
	return deleteFormat
}
