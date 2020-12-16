package v1

import (
	"net/http"
	"time"
	"xpertise-go/model"
	"xpertise-go/service"

	"github.com/gin-gonic/gin"
)

// Comment doc
// @description 评论
// @Tags branch
// @Param username formData string true "用户名"
// @Param paper_id formData string true "文献ID"
// @Param content formData string true "评论内容"
// @Success 200 {string} string "{"success": true, "message": "用户评论成功"}"
// @Router /branch/comment [post]
func Comment(c *gin.Context) {
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
