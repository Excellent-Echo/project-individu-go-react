package entity

type Questions struct {
	ID      uint64    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Title   string    `gorm:"size:255;not null" json:"title"`
	Content string    `gorm:"text;not null" json:"content"`
	UserID  uint32    `gorm:"not null" json:"user_id"`
	Tags    []Tags    `gorm:"many2many:question_tags" json:"tags"`
	Answers []Answers `gorm:"foreignKey:QuestionID" json:"answers"`
	// CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	// UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
	// DeletedAt time.Time `gorm:"type:datetime" json:"deleted_at"`
	// Likes     []Likes   `gorm:"foreignKey:QuestionID" json:"likes"`
}
