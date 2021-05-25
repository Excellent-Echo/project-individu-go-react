package entity

type Tags struct {
	ID  uint32 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Tag string `gorm:"size:255;not null" json:"tag"`
	// CreatedAt time.Time   `gorm:"type:datetime;not null;default:current_timestamp" json:"-"`
	Questions []Questions `gorm:"many2many:question_tags" json:"questions"`
}
