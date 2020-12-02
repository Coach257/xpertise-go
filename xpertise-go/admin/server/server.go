package server

import (
	"time"
	"xpertise-go/dao"
)

// ForbidAUser : Create a ForbidSpeak user.
func ForbidAUser(user *dao.User, duration *time.Duration) (err error) {
	item := dao.ForbidSpeak{
		UserID: user.UserID,
		StartTime: time.Now()
		Duration: duration
	}
	item.ExpireTime = item.StartTime + duration
	if err = dao.DB.Create(&item).Error; err != nil {
		return err
	}
	return
}

// ReportAComment : Create a ComReport
func ReportAComment(user *dao.User, comment *dao.Comment, content string) (err error) {
	item := dao.ComReport{
		UserID: user.UserID,
		CommentID: comment.CommentID
		Content: content,
		Status: 1
	}
	if err = dao.DB.Create(&item).Error; err != nil {
		return err
	}
	return
}
