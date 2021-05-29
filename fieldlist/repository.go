package fieldlist

import (
	"projectpenyewaanlapangan/entity"

	"gorm.io/gorm"
)

type Repository interface {
	// FindAll() ([]entity.FieldList, error)
	Create(fieldlist entity.FieldList) (entity.FieldList, error)
	FindByID(ID string) (entity.FieldList, error)
	FindAll() ([]entity.FieldListSportName, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) FindAll() ([]entity.FieldList, error) {
// 	var fieldlists []entity.FieldList

// 	if err := r.db.Find(&fieldlists).Error; err != nil {
// 		return fieldlists, err
// 	}

// 	return fieldlists, nil
// }

func (r *repository) Create(fieldlist entity.FieldList) (entity.FieldList, error) {
	if err := r.db.Create(&fieldlist).Error; err != nil {
		return fieldlist, err
	}

	return fieldlist, nil
}

func (r *repository) FindByID(ID string) (entity.FieldList, error) {
	var fieldlist entity.FieldList

	if err := r.db.Where("id = ?", ID).Find(&fieldlist).Error; err != nil {
		return fieldlist, err
	}

	return fieldlist, nil
}

func (e *repository) FindAll() ([]entity.FieldListSportName, error) {
	var FieldLists []entity.FieldListSportName

	if err := e.db.Model(entity.FieldList{}).Select("field_lists.id, field_lists.field_name, field_lists.field_image, field_lists.rent_price, field_lists.created_at, field_lists.updated_at, sport_lists.sport_name").Joins("left join sport_lists on sport_lists.id = field_lists.sport_id").Scan(&FieldLists).Error; err != nil {
		return FieldLists, err
	}

	return FieldLists, nil
}

// db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
