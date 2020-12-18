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
// @Param comment_id formData string true "评论ID"
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

// GiveALikeOrDislike doc
// @description 点赞或点踩
// @Tags branch
// @Param comment_id formData string true "评论ID"
// @Param user_id formData string true "用户ID"
// @Param method formData string true "对评论的操作选择，1为点赞，2为点踩"
// @Success 200 {string} string "{"success": true, "message": "用户操作成功"}"
// @Router /branch/comment/give_a_like_or_dislike [post]
func GiveALikeOrDislike(c *gin.Context) {
	var err error
	var comment model.Comment
	commentID, _ := strconv.ParseUint(c.Request.FormValue("comment_id"), 0, 64)
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	method, _ := strconv.ParseUint(c.Request.FormValue("method"), 0, 64)
	comment, commentNotFound := service.QueryAComment(commentID)
	commentLike, commentLikeNotFound := service.QueryAnItemFromCommentLike(commentID, userID)
	if commentNotFound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "没有找到相关的评论",
		})
		return
	}
	if commentLikeNotFound { // 用户没点过赞或踩，创建一条，并修改评论项的like/dislike
		err = service.CreateACommentLike(userID, &comment, method)
	} else if commentLike.LikeOrDislike && method == 2 { // 用户点过赞，需要进行转换，并修改评论表中的like/dislike
		err = service.TransferBetweenLikeAndDislike(&commentLike, &comment)
	} else if !commentLike.LikeOrDislike && method == 1 { // 用户点过踩，需要进行转换，并修改评论表中的like/dislike
		err = service.TransferBetweenLikeAndDislike(&commentLike, &comment)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "不能重复点赞或点踩",
		})
		return
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

// AuthorConnection doc
// @description 返回作者关系图数据
// @Tags branch
// @Param author_id formData string true "作者ID"
// @Param author_name formData string true "作者名字"
// @Success 200 {string} string "{"success": true, "message": {"rootID":"","nodes":[{"id":"","text":""}],"links":[{"from":"","to":"","text":""}]}}}"
// @Router /branch/graph/author_connection [post]
func AuthorConnection(c *gin.Context) {
	authorID := c.Request.FormValue("author_id")
	authorName := c.Request.FormValue("author_name")
	connection, notFound := service.GetFa(authorID)
	var nodes []model.Node
	var links []model.Link
	var ok bool
	if notFound {
		node := model.Node{Id: authorID, Text: authorName}
		nodes = append(nodes, node)
		linkType := model.LinkType{RootID: authorID, Nodes: nodes, Links: links}
		c.JSON(http.StatusOK, gin.H{"success": true, "message": linkType})
	} else {
		var authorMap map[string]bool
		authorMap = make(map[string]bool)
		connectionGraph := service.GetAuthorConnectionGraph(connection.FatherID)
		for _, v := range connectionGraph {
			_, ok = authorMap[v.Author1ID]
			if !ok {
				node := model.Node{Id: v.Author1ID, Text: v.Author1Name}
				nodes = append(nodes, node)
				authorMap[v.Author1ID] = true
			}
			_, ok = authorMap[v.Author2ID]
			if !ok {
				node := model.Node{Id: v.Author2ID, Text: v.Author2Name}
				nodes = append(nodes, node)
				authorMap[v.Author2ID] = true
			}
			link := model.Link{From: v.Author1ID, To: v.Author2ID, Text: v.PaperTitle}
			links = append(links, link)

		}
		linkType := model.LinkType{RootID: authorID, Nodes: nodes, Links: links}
		c.JSON(http.StatusOK, gin.H{"success": true, "message": linkType})
	}
	return
}

// ListAllComments doc
// @description 列出某条文献的全部评论
// @Tags branch
// @Param paper_id formData string true "文献ID"
// @Success 200 {string} string "{"success": true, "message": "操作成功", "data": "某文献的所有评论"}"
// @Router /branch/comment/list_all_comments [post]
func ListAllComments(c *gin.Context) {
	paperID := c.Request.FormValue("paper_id")
	comments := service.QueryAllComments(paperID)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "操作成功", "data": comments})
	return
}

func GetNodesFromReferences(references []model.PaperReference) []model.Node {
	set := make(map[model.PaperReference]bool)
	res := []model.Node{}
	for _, refer := range references {
		if !set[refer] {
			set[refer] = true
			tmp := model.Node{
				Id:   refer.ReferenceID,
				Text: refer.ReferencePaperTitle,
			}
			res = append(res, tmp)
		}
	}
	return res
}

func GetLinksFromReferences(references []model.PaperReference, level string) []model.Link {
	set := make(map[model.PaperReference]bool)
	var res []model.Link
	for _, refer := range references {
		if !set[refer] {
			set[refer] = true
			tmp := model.Link{
				From: refer.ReferenceID,
				To:   refer.ReferencePaperTitle,
				Text: level,
			}
			res = append(res, tmp)
		}
	}
	return res
}

// GetThreeLevelReferences doc
// @description 列出某条文献的三级参考文献
// @Tags branch
// @Param paper_id formData string true "文献ID"
// @Success 200 {string} string "{"success": true, "message": "操作成功", "data": "某文献的2级参考文献"}"
// @Router /branch/reference [post]
func GetThreeLevelReferences(c *gin.Context) {
	paperID := c.Request.FormValue("paper_id")
	directRefers := service.QueryAllReferences(paperID)
	nodes := GetNodesFromReferences(directRefers)
	links := GetLinksFromReferences(directRefers, "1")
	var secondRefers = []model.PaperReference{}
	for _, refer := range secondRefers {
		tmpRefers := service.QueryAllReferences(refer.ReferenceID)
		secondRefers = append(secondRefers, tmpRefers...)
	}
	nodes = append(nodes, GetNodesFromReferences(directRefers)...)
	links = append(links, GetLinksFromReferences(directRefers, "2")...)
	data := model.LinkType{
		RootID: paperID,
		Nodes:  nodes,
		Links:  links,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "操作成功", "data": data})
	return
}
