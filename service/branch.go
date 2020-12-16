package service

import (
	"xpertise-go/global"
	"xpertise-go/model"
)

func CreateAComment(comment *model.Comment) (err error) {
	if err = global.DB.Create(&comment).Error; err != nil {
		return err
	}
	return
}
