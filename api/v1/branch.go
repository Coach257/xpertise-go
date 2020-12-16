package v1

import (
	"net/http"
	"strconv"
	"time"
	"xpertise-go/model"
	"xpertise-go/service"

	"github.com/gin-gonic/gin"
)

// CreateAComment doc
// @description 创建一条评论
// @Tags branch
// @Param username formData string true "用户名"
// @Param paper_id formData string true "文献ID"
// @Param content formData string true "评论内容"
// @Success 200 {string} string "{"success": true, "message": "用户评论成功"}"
// @Router /branch/comment/create [post]
func CreateAComment(c *gin.Context) {
	username := c.Request.FormValue("username")
	paperID := c.Request.FormValue("paper_id")
	content := c.Request.FormValue("content")

	comment := model.Comment{Username: username, PaperID: paperID, CommentTime: time.Now(), Content: content}
	err := service.CreateAComment(&comment)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "评论失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "用户评论成功"})
	return
}

// OperateComment doc
// @description 操作评论
// @Tags branch
// @Param username formData string true "评论ID"
// @Param method formData string true "对评论的操作方法，1为置顶，2为取消置顶，3为删除"
// @Success 200 {string} string "{"success": true, "message": "操作成功"}"
// @Router /branch/comment/operate [post]
func OperateComment(c *gin.Context) {
	commentID, _ := strconv.ParseUint(c.Request.FormValue("comment_id"), 0, 64)
	method, _ := strconv.ParseUint(c.Request.FormValue("method"), 0, 64)
	var err error
	switch method {
	case 1:
		err = service.PutCommentToTop(commentID)
	case 2:
		err = service.CancelCommentToTop(commentID)
	case 3:
		err = service.DeleteAComment(commentID)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "操作成功"})
	return
}
