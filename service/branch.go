package service

import (
	"xpertise-go/global"
	"xpertise-go/model"

	"github.com/jinzhu/gorm"
)

func CreateAComment(comment *model.Comment) (err error) {
	if err = global.DB.Create(&comment).Error; err != nil {
		return err
	}
	return
}

// 删除某条评论
func DeleteAComment(CommentID uint64) (err error) {
	var comment model.Comment
	notFound := global.DB.First(&comment, CommentID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	err = global.DB.Delete(&comment).Error
	return err
}

// 置顶某条评论
func PutCommentToTop(commentID uint64) (err error) {
	var comment model.Comment
	notFound := global.DB.First(&comment, commentID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	comment.OnTop = true
	err = global.DB.Save(comment).Error
	return err
}

// 取消置顶某条评论
func CancelCommentToTop(commentID uint64) (err error) {
	var comment model.Comment
	notFound := global.DB.First(&comment, commentID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	comment.OnTop = false
	err = global.DB.Save(comment).Error
	return err
}
