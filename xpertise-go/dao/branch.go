package dao

import "time"

type Comment struct {
	//gorm.Model
	ID        uint64 `gorm:"primary_key;column:com_id"`
	UserID    uint64 `gorm:"column:user_id"`
	DocID     uint64 `gorm:"column:doc_id"`
	CreatedAt time.Time
	Content   string `gorm:"column:content;size:255"` // string默认长度为255, 使用这种tag重设。

	//PreCom  uint64 `gorm:"column:pre_comment"`
	Like    uint64
	DisLike uint64
}

// 在portal里重复定义了Document，这边的先注视掉
// type Document struct {
// 	//gorm.Model
// 	ID          uint64    `gorm:"primary_key;column:doc_id"`
// 	Comments    []Comment `gorm:"ForeignKey:DocID"`
// 	AuthorList  string    `gorm:"column:author_list"`
// 	TypeList    string    `gorm:"column:type_list"`
// 	Abstract    string    `gorm:"column:abstract;size:255"`
// 	CiteList    string    `gorm:"column:cite_list"`
// 	PublishTime time.Time
// 	Source      string `gorm:"column:source;size:255"`
// 	Content     string `gorm:"column:content;size:255"` // string默认长度为255, 使用这种tag重设。

// }

type LikeDislikeRecord struct {
	ComID       uint64 `gorm:"column:com_id"`
	UserID      uint64 `gorm:"column:user_id"`
	IsLikeOrDis bool
}
