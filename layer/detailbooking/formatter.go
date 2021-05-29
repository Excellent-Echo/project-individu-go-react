package detailbooking

import "project-individu-go-react/entity"

type BookingDetailFormat struct {
	ID         int `json:"id"`
	BookingID  int `json:"booking_id"`
	PsikologID int `json:"psikolog_id"`
}

func FormatBookingDetails(bookingDet entity.BookingDetail) BookingDetailFormat {
	var formatBookingDetails = BookingDetailFormat{
		ID:         bookingDet.ID,
		BookingID:  bookingDet.BookingID,
		PsikologID: bookingDet.PsikologID,
	}
	return formatBookingDetails
}
