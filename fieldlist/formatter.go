package fieldlist

import (
	"projectpenyewaanlapangan/entity"
	"time"
)

type FieldListFormat struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	FieldName  string `json:"field_name"`
	FieldImage string `json:"field_image"`
	RentPrice  int    `json:"rent_price"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatFieldList(fieldlist entity.FieldList) FieldListFormat {
	var formatFieldList = FieldListFormat{
		ID:         fieldlist.ID,
		FieldName:  fieldlist.FieldName,
		FieldImage: fieldlist.FieldImage,
		RentPrice:  fieldlist.RentPrice,
	}

	return formatFieldList
}
