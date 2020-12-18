package model

import (
	"time"
)

type Paper struct {
	PaperID          string `gorm:"type:varchar(30);primary_key;" json:"paper_id"`
	Title            string `gorm:"type:varchar(400);not null;" json:"title"`
	PaperPublishYear string `gorm:"type:varchar(5)" json:"paper_publish_year"`
	ConferenceID     string `gorm:"type:varchar(10)" json:"conference_id"`
}

type Affiliation struct {
	AffiliationID   string `gorm:"type:varchar(10);primary_key;" json:"affiliation_id"`
	AffiliationName string `gorm:"type:varchar(150)" json:"affiliation_name"`
}

type Author struct {
	AuthorID   string `gorm:"type:varchar(30);primary_key;" json:"author_id"`
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
	PaperID             string `gorm:"type:varchar(10);" json:"paper_id"`
	PaperTitle          string `gorm:"type:varchar(400);" json:"paper_title"`
	ReferenceID         string `gorm:"type:varchar(10);" json:"reference_id"`
	ReferencePaperTitle string `gorm:"type:varchar(400);" json:"reference_paper_title"`
}

type Connection struct {
	Author1ID   string `gorm:"type:varchar(10)" json:"author1_id"`
	Author1Name string `gorm:"type:varchar(100)" json:"author1_name"`
	Author2ID   string `gorm:"type:varchar(10)" json:"author2_id"`
	Author2Name string `gorm:"type:varchar(100)" json:"author2_name"`
	FatherID    string `gorm:"type:varchar(10)" json:"fa_id"`
	PaperID     string `gorm:"type:varchar(10)" json:"paper_id"`
	PaperTitle  string `gorm:"type:varchar(400);" json:"paper_title"`
}

type Comment struct {
	//gorm.Model
	CommentID   uint64    `gorm:"primary_key;" json:"comment_id"`
	Username    string    `gorm:"size:15" json:"username"`
	PaperID     string    `gorm:"size:30" json:"paper_id"`
	CommentTime time.Time `json:"comment_time"`
	Content     string    `gorm:"size:255" json:"content"`
	OnTop       bool      `gorm:"default:false" json:"on_top"`
	Like        uint64    `json:"like"`
	Dislike     uint64    `json:"dislike"`
}

type CommentLike struct {
	CommentID     uint64 `gorm:"primary_key;" json:"comment_id"`
	UserID        uint64 `gorm:"primary_key;" json:"user_id"`
	LikeOrDislike bool   `json:"like_or_dislike"`
}

type LinkType struct {
	RootID string `json:"rootID"`
	Nodes  []Node `json:"nodes"`
	Links  []Link `json:"links"`
}
type Node struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}
type Link struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}
