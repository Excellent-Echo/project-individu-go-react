package fieldlist

import "projectpenyewaanlapangan/entity"

func FormatFieldList(fieldlist entity.FieldList) entity.FieldList {
	var FormatFieldList = entity.FieldList{
		ID:         fieldlist.ID,
		FieldName:  fieldlist.FieldName,
		FieldImage: fieldlist.FieldImage,
		RentPrice:  fieldlist.RentPrice,
		SportID:    fieldlist.SportID,
	}

	return FormatFieldList
}

func FormatFieldListSportName(fieldlist entity.FieldListSportName) entity.FieldListSportName {
	var FormatFieldListSportName = entity.FieldListSportName{
		ID:         fieldlist.ID,
		FieldName:  fieldlist.FieldName,
		FieldImage: fieldlist.FieldImage,
		RentPrice:  fieldlist.RentPrice,
		SportID:    fieldlist.SportID,
		SportName:  fieldlist.SportName,
	}

	return FormatFieldListSportName
}
