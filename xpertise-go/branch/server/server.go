package server

import (
	"xpertise-go/dao"
)

// CreatAComment 创建一条评论
func CreateAComment(comm *dao.Comment) (err error) {

	if err = dao.DB.Create(&comm).Error; err != nil {
		return err
	}
	return
}

func DeleteACommentByID(CommentID uint64) {
	dao.DB.Where("com_id = ?", CommentID).Delete(&dao.Comment{})
	return
}

func CommentCheck(useridInt uint64, comidInt uint64) (blank bool, likeOrDislike bool) {

	result := new(dao.LikeDislikeRecord)
	dao.DB.Where(dao.LikeDislikeRecord{ComID: comidInt, UserID: useridInt}).First(&result)
	//fmt.Println(result)
	if result.ComID != 0 {
		blank = true
		likeOrDislike = result.IsLikeOrDis
		//fmt.Println(likeOrDislike)
		return blank, likeOrDislike
	}
	blank = false
	return blank, likeOrDislike
}

func CreateALike(like *dao.LikeDislikeRecord) (err error) {

	if err = dao.DB.Create(&like).Error; err != nil {
		return err
	}
	return
}
func CreateADisLike(dislike *dao.LikeDislikeRecord) (err error) {

	if err = dao.DB.Create(&dislike).Error; err != nil {
		return err
	}
	return
}

func DisLikeTolike(useridInt uint64, comidInt uint64) {
	dao.DB.Model(&dao.LikeDislikeRecord{}).Where(dao.LikeDislikeRecord{ComID: comidInt, UserID: useridInt}).Update("IsLikeOrDis", true)
	return
}
func LikeToDislike(useridInt uint64, comidInt uint64) {
	dao.DB.Model(&dao.LikeDislikeRecord{}).Where(dao.LikeDislikeRecord{ComID: comidInt, UserID: useridInt}).Update("IsLikeOrDis", false)
	return
}
func DeleteThumbUp(userID uint64, commentID uint64) {
	dao.DB.Where("user_id = ? AND com_id = ?", userID, commentID).Delete(&dao.LikeDislikeRecord{})
	return
}

func DeleteThumbDown(userID uint64, commentID uint64) {
	dao.DB.Where("user_id = ? AND com_id = ?", userID, commentID).Delete(&dao.LikeDislikeRecord{})
	return
}

func LikeAdd(comidInt uint64) {
	result := new(dao.Comment)
	dao.DB.Where(dao.Comment{ComID: comidInt}).First(&result)
	dao.DB.Model(&dao.Comment{}).Where(dao.Comment{ComID: comidInt}).Update("Like", result.Like+1)
	return

}

func DisLikeAdd(comidInt uint64) {
	result := new(dao.Comment)
	dao.DB.Where(dao.Comment{ComID: comidInt}).First(&result)
	dao.DB.Model(&dao.Comment{}).Where(dao.Comment{ComID: comidInt}).Update("DisLike", result.DisLike+1)
	return

}

func LikeDec(comidInt uint64) {
	result := new(dao.Comment)
	dao.DB.Where(dao.Comment{ComID: comidInt}).First(&result)
	dao.DB.Model(&dao.Comment{}).Where(dao.Comment{ComID: comidInt}).Update("Like", result.Like-1)
	return

}

func DisLikeDec(comidInt uint64) {
	result := new(dao.Comment)
	dao.DB.Where(dao.Comment{ComID: comidInt}).First(&result)
	dao.DB.Model(&dao.Comment{}).Where(dao.Comment{ComID: comidInt}).Update("DisLike", result.DisLike-1)
	return

}

func QueryComThumbUp(comidInt uint64) (thumb uint64) {
	result := new(dao.Comment)
	dao.DB.Where(dao.Comment{ComID: comidInt}).First(&result)
	thumb = result.Like
	return thumb
}
func QueryComThumbDown(comidInt uint64) (thumb uint64) {
	result := new(dao.Comment)
	dao.DB.Where(dao.Comment{ComID: comidInt}).First(&result)
	thumb = result.DisLike
	return thumb
}
