package server

import (
	"time"
	"xpertise-go/dao"
)

// QueryAComReportByID query a comreport
func QueryAComReportByID(comReportID uint64) (comReport dao.ComReport, NotFound bool) {
	NotFound = dao.DB.First(&comReport, comReportID).RecordNotFound()
	return comReport, NotFound
}

// UpdateAComReportByStatus update a comreport's status
func UpdateAComReportByStatus(comReport *dao.ComReport, status string) {
	dao.DB.Model(&comReport).Update("Status", status)
}

// ForbidAUser : Create a ForbidSpeak user.
func ForbidAUser(user *dao.User, duration *time.Duration) (err error) {
	item := dao.ForbidSpeak{
		UserID:    user.UserID,
		StartTime: time.Now(),
		Duration:  *duration,
	}
	item.ExpireTime = item.StartTime.Add(*duration)
	if err = dao.DB.Create(&item).Error; err != nil {
		return err
	}
	return
}

// ReportAComment : Create a ComReport
func ReportAComment(user *dao.User, comment *dao.Comment, content string) (err error) {
	item := dao.ComReport{
		UserID:    user.UserID,
		CommentID: comment.ComID,
		Content:   content,
		Status:    "Wait",
	}
	if err = dao.DB.Create(&item).Error; err != nil {
		return err
	}
	return
}
