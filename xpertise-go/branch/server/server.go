package server

import "xpertise-go/branch/dao"

// CreatAComment 创建一条评论
func CreateAComment(comm *dao.Comment) (err error) {
	if err = dao.DB.Create(&comm).Error; err != nil {
		return err
	}
	return
}
