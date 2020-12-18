package model

import "time"

// AuthorizationRequest 申请认证
type AuthorizationRequest struct {
	AuthReqID    uint64    `gorm:"primary_key; not null;" json:"authreq_id"`
	UserID       uint64    `gorm:"not null;" json:"user_id"`
	AuthorID     string    `gorm:"type:varchar(30)" json:"author_id"`
	CitizenID    string    `gorm:"type:varchar(20)" json:"citizen_id"`
	Organization string    `gorm:"type:varchar(256)" json:"organization"`
	ReqTime      time.Time `gorm:"type:datetime" json:"request_time"`
	Status       string    `gorm:"type:varchar(10)" json:"status"`
}

// // ForbidSpeak 禁言用户
// type ForbidSpeak struct {
// 	ForbidID   uint64        `gorm:"primary_key; not null"`
// 	UserID     uint64        `gorm:"ForeignKey:UserID;"`
// 	StartTime  time.Time     `gorm:"not null" json:"start_time"`
// 	Duration   time.Duration `gorm:"not null" json:"duration"`
// 	ExpireTime time.Time     `gorm:"not null" json:"expire_time"`
// 	Reason     string        `gorm:"size:255" json:"reason"`
// 	//User_type Profile `gorm:"ForeignKey:UserRefer"` 指定外键
// }
