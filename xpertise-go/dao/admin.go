package dao

import "time"

// ForbidSpeak 禁言用户
type ForbidSpeak struct {
	ForbidID   uint64 `gorm:"primary_key; not null"`
	UserID     uint64 `gorm:"column:user_id"`
	StartTime  time.Time
	Duration   time.Duration
	ExpireTime time.Time
	Reason     string `gorm:"column:reason;size:255"`
	//User_type Profile `gorm:"ForeignKey:UserRefer"` 指定外键
}

// ComReport 举报评论
type ComReport struct {
	ReportID  uint64 `gorm:"primary_key; not null"`
	UserID    uint64 `gorm:"column:user_id; not null"`
	CommentID uint64 `gorm:"column:comment_id; not null"`
	Content   string `gorm:"column:content;size:255"`
	Status    string `gorm:"column:status; not null"`
}

// TransferPatent dd
type TransferPatent struct {
	TransPatID   uint64    `gorm:"primary_key;" json:"trans_pat_id"`
	SenderID     uint64    `gorm:"ForeignKey:OrgID;not null;" json:"sender_id"`
	ReceiverID   uint64    `gorm:"ForeignKey:OrgID;not null;" json:"receive_id"`
	PatentID     uint64    `gorm:"ForeignKey:PatentID;not null;" json:"patent_id"`
	TransferTime time.Time `json:"transfer_time"`
	Status       uint64    `json:"staus"`
}
