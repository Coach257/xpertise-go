package service

import (
	"xpertise-go/global"
	"xpertise-go/model"

	"github.com/jinzhu/gorm"
)

// 创建一条认证申请
func CreateAnAuthorizationRequest(userID uint64, authorID string) (err error) {
	authreq := model.AuthorizationRequest{
		UserID:   userID,
		AuthorID: authorID,
	}
	if err = global.DB.Create(&authreq).Error; err != nil {
		return err
	}
	return
}

// 获取一条认证申请
func QueryAnAuthorizationRequest(authreqID uint64) (authreq model.AuthorizationRequest, err error) {
	notFound := global.DB.First(&authreq, authreqID).RecordNotFound()
	if notFound {
		return authreq, gorm.ErrRecordNotFound
	}
	return authreq, err
}

// DeleteAuthorizationRequest 删除一条认证申请
func DeleteAuthorizationRequest(authreqID uint64) (err error) {
	var authreq model.AuthorizationRequest
	notFound := global.DB.First(&authreq, authreqID).RecordNotFound()
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
