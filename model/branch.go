package model

import (
	"time"
)

type Comment struct {
	//gorm.Model
	CommentID   uint64    `gorm:"primary_key;" json:"comment_id"`
	Username    string    `gorm:"size:15" json:"username"`
	PaperID     string    `gorm:"size:10" json:"paper_id"`
	CommentTime time.Time `json:"comment_time"`
	Content     string    `gorm:"size:255" json:"content"`
	OnTop       bool      `gorm:"default:false" json:"on_top"`
}
