package bookinglist

import (
	"time"
)

// type BookingFormat struct {
// 	ID        int       `json:"id"`
// 	Date      time.Time `json:"date"`
// 	FirstName string    `json:"first_name"`
// 	LastName  string    `json:"last_name"`
// 	User_id   int       `json:"userID"`
// }

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

// func FormatBooking(booking entity.BookingList) BookingFormat {
// 	var formatBooking = BookingFormat{
// 		ID:        booking.ID,
// 		Date:      booking.Date,
// 		FirstName: booking.FirstName,
// 		LastName:  booking.LastName,
// 		User_id:   booking.User_id,
// 	}
// 	return formatBooking
// }

func FormatDeleteUser(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
