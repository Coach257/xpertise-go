package service

import (
	"xpertise-go/global"
	"xpertise-go/model"

	"github.com/jinzhu/gorm"
)

// CreateAnAuthorizationRequest 创建一条认证申请
func CreateAnAuthorizationRequest(userID uint64) (err error) {
	authreq := model.AuthorizationRequest{
		UserID:   userID,
		AuthorID: "",
		Status:   "TODO",
	}
	if err = global.DB.Create(&authreq).Error; err != nil {
		return err
	}
	return
}

// QueryAnAuthorizationRequest 获取一条认证申请
func QueryAnAuthorizationRequest(authreqID uint64) (authreq model.AuthorizationRequest, err error) {
	notFound := global.DB.First(&authreq, authreqID).RecordNotFound()
	if notFound {
		return authreq, gorm.ErrRecordNotFound
	}
	return authreq, err
}

// QueryAuthorizationRequestsByUserID 获取某个用户的所有认证申请
func QueryAuthorizationRequestsByUserID(userID uint64) (authreqs []model.AuthorizationRequest, err error) {
	notFound := global.DB.Find(&authreqs, userID).RecordNotFound()
	if notFound {
		return authreqs, gorm.ErrRecordNotFound
	}
	return authreqs, err
}

// UpdateAnAuthorizationRequest 更新一条认证申请
func UpdateAnAuthorizationRequest(authreqID uint64, status string, authorID string) (err error) {
	var authreq model.AuthorizationRequest
	notFound := global.DB.First(&authreq, authreqID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	authreq.Status = status
	if status == "Accepted" {
		authreq.AuthorID = authorID
	}
	err = global.DB.Update(authreq).Error
	return err
}

// DeleteAnAuthorizationRequest 删除一条认证申请
func DeleteAnAuthorizationRequest(authreqID uint64, userID uint64) (err error) {
	var authreq model.AuthorizationRequest
	notFound := global.DB.First(&authreq, authreqID, userID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	err = global.DB.Delete(authreq).Error
	return err
}

// QueryAllAuthorizationRequest 列出所有认证申请
func QueryAllAuthorizationRequest() (authreqs []model.AuthorizationRequest) {
	global.DB.Find(&authreqs)
	return authreqs
}
