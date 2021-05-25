package fieldlist

import "projectpenyewaanlapangan/entity"

func FormatFieldList(fieldlist entity.FieldList) entity.FieldList {
	var FormatFieldList = entity.FieldList{
		ID:         fieldlist.ID,
		FieldName:  fieldlist.FieldName,
		FieldImage: fieldlist.FieldImage,
		SportID:    fieldlist.SportID,
	}

	return FormatFieldList
}
