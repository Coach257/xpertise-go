package model

type Paper struct {
	PaperID          string `gorm:"type:varchar(10);primary_key;" json:"paper_id"`
	Title            string `gorm:"type:varchar(400);not null;" json:"title"`
	PaperPublishYear string `gorm:"type:varchar(5)" json:"paper_publish_year"`
	ConferenceID     string `gorm:"type:varchar(10)" json:"conference_id"`
}
