package service

import (
	"xpertise-go/global"
	"xpertise-go/model"

	"github.com/jinzhu/gorm"
)

// CreateAuthorizationRequest 创建一条认证申请
func CreateAuthorizationRequest(userID uint64, authorID string) (err error) {
	authreq := model.AuthorizationRequest{
		UserID:   userID,
		AuthorID: authorID,
	}
	if err = global.DB.Create(&authreq).Error; err != nil {
		return err
	}
	return
}

// QueryAuthorizationRequest 获取一条认证申请
func QueryAAuthorizationRequest(authreqID uint64) (authreq model.AuthorizationRequest, err error) {
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

// QueryAllAuthorizationRequest 列出某专栏中的所有内容
func QueryAllAuthorizationRequest() (authreqs []model.AuthorizationRequest) {
	global.DB.Find(&authreqs)
	return authreqs
}
