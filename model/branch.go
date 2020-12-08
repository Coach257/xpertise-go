package model

import (
	"time"
)

type Comment struct {
	//gorm.Model
	CommentID  uint64    `gorm:"primary_key" json:"comment_id"`
	UserID     uint64    `gorm:"not null;ForeignKey:UserID" json:"user_id"`
	DocumentID uint64    `gorm:"not null;ForeignKey:DocumentID" json:"document_id"` // 外键
	CreatedAt  time.Time `json:"created_at"`
	Content    string    `gorm:"column:content;size:255" json:"content"` // string默认长度为255, 使用这种tag重设。
	Like       uint64    `gorm:"default:0" json:"like"`
	DisLike    uint64    `gorm:"default:0" json:"dislike"`
}
