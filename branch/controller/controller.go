package controller

import (
	"net/http"
	"strconv"
	"xpertise-go/branch/server"
	"xpertise-go/dao"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Xpertise Scholar",
	})

}

func InputBlankCheck(content string) (bool, string) {
	if content == "" {
		return true, "评论不能为空"
	}

	return false, ""
}
func CreateAComment(c *gin.Context) {

	userid := c.Request.FormValue("UserID")
	useridInt, _ := strconv.ParseUint(userid, 0, 64)
	docid := c.Request.FormValue("DocID")
	docidInt, _ := strconv.ParseUint(docid, 0, 64)
	content := c.Request.FormValue("Content")

	if blank, message := InputBlankCheck(content); blank {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": message,
		})
		return

	}
	comment := dao.Comment{UserID: useridInt, DocID: docidInt, Content: content}
	server.CreateAComment(&comment)
	c.JSON(200, gin.H{"success": true, "message": "评论创建成功"})
}

func DeleteACommentByID(c *gin.Context) {
	cid, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	server.DeleteACommentByID(cid)
	c.JSON(200, gin.H{"message": "success"})
}
func AddLike(c *gin.Context) {

	userid := c.Request.FormValue("UserID")
	useridInt, _ := strconv.ParseUint(userid, 0, 64)
	comid := c.Request.FormValue("ComID")
	comidInt, _ := strconv.ParseUint(comid, 0, 64)
	if blank, likeOrDislike := server.CommentCheck(useridInt, comidInt); blank {
		if likeOrDislike == true {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "已经点赞",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "已踩，改为点赞",
		})
		server.DisLikeTolike(useridInt, comidInt)
		server.LikeAdd(comidInt)
		server.DisLikeDec(comidInt)
		return

	}

	like := dao.LikeDislikeRecord{UserID: useridInt, ComID: comidInt, IsLikeOrDis: true}
	server.CreateALike(&like)
	server.LikeAdd(comidInt)
	c.JSON(200, gin.H{"success": true, "message": "点赞成功"})
}

func RevertAddLike(c *gin.Context) {

	userid := c.Request.FormValue("UserID")
	useridInt, _ := strconv.ParseUint(userid, 0, 64)
	comid := c.Request.FormValue("ComID")
	comidInt, _ := strconv.ParseUint(comid, 0, 64)

	server.DeleteThumbUp(useridInt, comidInt)
	server.LikeDec(comidInt)
	c.JSON(200, gin.H{"message": "赞取消成功"})
}

func AddDisLike(c *gin.Context) {

	userid := c.Request.FormValue("UserID")
	useridInt, _ := strconv.ParseUint(userid, 0, 64)
	comid := c.Request.FormValue("ComID")
	comidInt, _ := strconv.ParseUint(comid, 0, 64)
	if blank, likeOrDislike := server.CommentCheck(useridInt, comidInt); blank {
		if likeOrDislike == false {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "已经踩了",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "已点赞，改为踩",
		})

		server.LikeToDislike(useridInt, comidInt)
		server.LikeDec(comidInt)
		server.DisLikeAdd(comidInt)
		return

	}

	dislike := dao.LikeDislikeRecord{UserID: useridInt, ComID: comidInt, IsLikeOrDis: false}
	server.CreateADisLike(&dislike)
	server.DisLikeAdd(comidInt)
	c.JSON(200, gin.H{"success": true, "message": "踩成功"})
}

func RevertAddDisLike(c *gin.Context) {

	userid := c.Request.FormValue("UserID")
	useridInt, _ := strconv.ParseUint(userid, 0, 64)
	comid := c.Request.FormValue("ComID")
	comidInt, _ := strconv.ParseUint(comid, 0, 64)

	server.DeleteThumbDown(useridInt, comidInt)
	server.DisLikeDec(comidInt)
	c.JSON(200, gin.H{"message": "踩取消成功"})
}

func QueryThumbUp(c *gin.Context) {
	cid, _ := strconv.ParseUint(c.Request.FormValue("id"), 0, 64)
	thumbUp := server.QueryComThumbUp(cid)
	c.IndentedJSON(200, thumbUp)
}

func QueryThumbDown(c *gin.Context) {
	cid, _ := strconv.ParseUint(c.Request.FormValue("id"), 0, 64)
	thumbDown := server.QueryComThumbDown(cid)
	c.IndentedJSON(200, thumbDown)
}
