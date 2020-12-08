package model

import (
	"time"
	"xpertise-go/utils"
)

type Document struct {
	DocumentID  uint64        `gorm:"primary_key;" json:"document_id"`
	Title       string        `gorm:"size:250;not null;" json:"title" binding:"required"`
	AuthorList  utils.StrList `gorm:"type:varchar(300)" json:"author_list"` // 作者列表
	TypeList    utils.StrList `gorm:"type:varchar(300)" json:"type_list"`   // 领域列表
	Abstract    string        `gorm:"size:5000" json:"abstract"`
	CiteList    utils.StrList `gorm:"type:varchar(500)" json:"cite_list"` //引用列表
	PublishTime time.Time     `json:"publish_time"`
	Source      string        `gorm:"size:250" json:"source"`    // 所属期刊
	Content     string        `gorm:"size:1000" json:"original"` // 原文
	Comments    []Comment     `gorm:"ForeignKey:DocumentID" json:"comments"`
}
