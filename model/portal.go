package model

type Paper struct {
	PaperID          string `gorm:"type:varchar(10);primary_key;" json:"paper_id"`
	Title            string `gorm:"type:varchar(400);not null;" json:"title"`
	PaperPublishYear string `gorm:"type:varchar(5)" json:"paper_publish_year"`
	ConferenceID     string `gorm:"type:varchar(10)" json:"conference_id"`
}

type Affiliation struct {
	AffiliationID   string `gorm:"type:varchar(10);primary_key;" json:"affiliation_id"`
	AffiliationName string `gorm:"type:varchar(150)" json:"affiliation_name"`
}

type Author struct {
	AuthorID   string `gorm:"type:varchar(10);primary_key;" json:"author_id"`
	AuthorName string `gorm:"type:varchar(100)" json:"author_name"`
}

type Conference struct {
	ConferenceID   string `gorm:"type:varchar(10);primary_key;" json:"conference_id"`
	ConferenceName string `gorm:"type:varchar(10)" json:"conference_name"`
}

type PaperAuthorAffiliation struct {
	PaperID        string `gorm:"type:varchar(10);primary_key;" json:"paper_id"`
	AuthorID       string `gorm:"type:varchar(10)" json:"author_id"`
	AffiliationID  string `gorm:"type:varchar(10)" json:"affiliation_id"`
	AuthorSequence string `gorm:"type:varchar(3);primary_key;" json:"author_sequence"`
}

type PaperReference struct {
	PaperID     string `gorm:"type:varchar(10);primary_key" json:"paper_id"`
	ReferenceID string `gorm:"type:varchar(10);primary_key" json:"reference_id"`
}

type SpecialColumn struct {
	ColumnID   uint64 `gorm:"primary_key; not null;" json:"column_id"`
	AuthorID   string `gorm:"type:varchar(10);primary_key;" json:"author_id"`
	ColumnName string `gorm:"type:varchar(100)" json:"column_name"`
}

type ColumnPaper struct {
	ColumnID uint64 `gorm:"primary_key; not null;" json:"column_id"`
	PaperID  string `gorm:"type:varchar(10);primary_key; not null;" json:"paper_id"`
}
