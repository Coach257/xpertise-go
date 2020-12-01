package dao

import "time"

type ForbidSpeak struct {
	ForbidID     uint64 `gorm:"primary_key; not null"`
	UserID       uint64 `gorm:"column:user_id"`
	StartTime    time.Time
	Duration     time.Time
	ExpireTime   time.Time
	Reason       string `gorm:"column:reason;size:255"`
	//User_type Profile `gorm:"ForeignKey:UserRefer"` 指定外键
}

type ComReport struct {
	ReportID     uint64 `gorm:"primary_key; not null"`
	UserID       uint64 `gorm:"column:user_id; not null"`
	CommentID    uint64 `gorm:"column:comment_id; not null"`
	Content      string `gorm:"column:content;size:255"`
	status       uint64 `gorm:"column:status; not null"`
}
