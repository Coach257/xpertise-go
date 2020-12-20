package v1

import (
	"net/http"
	"strconv"
	"xpertise-go/service"

	"github.com/gin-gonic/gin"
)

// CreateSpecialColumn doc
// @description 创建一个专栏
// @Tags portal
// @Param author_id formData string true "作者ID"
// @Param column_name formData string true "专栏名字"
// @Success 200 {string} string "{"success": true, "message": "创建专栏成功"}"
// @Router /portal/column/create_column [post]
func CreateSpecialColumn(c *gin.Context) {
	authorID := c.Request.FormValue("author_id")
	columnName := c.Request.FormValue("column_name")
	if err := service.CreateAColumn(authorID, columnName); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "创建专栏成功"})
	return
}

// SearchSpecialColumn doc
// @description 返回某个作者的一个专栏
// @Tags portal
// @Param author_id formData string true "作者ID"
// @Success 200 {string} string "{"success": true, "message": "返回专栏成功"}"
// @Router /portal/column/searchcol [post]
func SearchSpecialColumn(c *gin.Context) {
	authorID := c.Request.FormValue("author_id")
	data := service.QueryAColumnByID(authorID)
	if data == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "不对"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "返回专栏成功", "data": data})
	return
}

// AddToColumn doc
// @description 添加某篇文章到专栏
// @Tags portal
// @Param paper_id formData string true "文献ID"
// @Param column_id formData string true "专栏ID"
// @Param paper_title formData string true "文献标题"
// @Success 200 {string} string "{"success":true, "message":"添加到专栏成功"}"
// @Router /portal/column/add_to_column [post]
func AddToColumn(c *gin.Context) {
	columnID, _ := strconv.ParseUint(c.Request.FormValue("column_id"), 0, 64)
	paperID := c.Request.FormValue("paper_id")
	paperTitle := c.Request.FormValue("paper_title")
	_, notFound := service.QueryItemFromColumnPaper(columnID, paperID)

	if !notFound {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "不能重复添加"})
		return
	}

	if err := service.AddPaperToColumn(columnID, paperID, paperTitle); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "添加到专栏成功"})
	return
}

// ListAllFromAColumn doc
// @description 获取某个专栏的所有内容
// @Tags portal
// @Param column_id formData string true "专栏ID"
// @Success 200 {string} string "{"success": true, "message": "查找成功", "data": "专栏中的所有论文ID"}"
// @Router /portal/column/list_all_from_column [post]
func ListAllFromAColumn(c *gin.Context) {
	columnID, _ := strconv.ParseUint(c.Request.FormValue("column_id"), 0, 64)
	columnPapers := service.QueryAllFromAColumn(columnID)

	if len(columnPapers) == 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "没查询到内容", "data": columnPapers})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查找成功", "data": columnPapers})
	return
}

// RemovePaperFromColumn doc
// @description 删除专栏中的某条论文
// @Tags portal
// @Param column_id formData string true "专栏ID"
// @Param paper_id formData string true "论文ID"
// @Success 200 {string} string "{"success": true, "message": "删除成功"}"
// @Router /portal/column/remove_from_column [post]
func RemovePaperFromColumn(c *gin.Context) {
	columnID, _ := strconv.ParseUint(c.Request.FormValue("column_id"), 0, 64)
	paperID := c.Request.FormValue("paper_id")

	if err := service.DeleteOnePaperFromAColumn(columnID, paperID); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "删除成功"})
	}

}

// SearchAuthor doc
// @description 查找作者是否存在
// @Tags portal
// @Param id formData string true "作者ID"
// @Success 200 {string} string "{"success": true, "message": "查询成功", "data": au}"
// @Router /portal/author [post]
func SearchAuthor(c *gin.Context) {
	//authorID, _ := strconv.ParseUint(c.Request.FormValue("id"), 0, 64)
	authorID := c.Request.FormValue("id")

	if au, notFound := service.QueryAnAuthorByID(authorID); notFound == true {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "cuowu"})

	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": au})
	}
	return
}

// CreateRecommend doc
// @description 创建一条推荐
// @Tags portal
// @Param author_id formData string true "作者ID"
// @Param author_name formData string true "作者名字"
// @Param paper_id formData string true "文献ID"
// @Param paper_title formData string true "文献名"
// @Param n_citation formData string true "引用次数"
// @Param hindex formData string true "h-index"
// @Param reason formData string true "推荐理由"
// @Success 200 {string} string "{"success": true, "message": "推荐成功"}"
// @Router /portal/recommend/create [post]
func CreateRecommend(c *gin.Context) {
	authorID := c.Request.FormValue("author_id")
	authorName := c.Request.FormValue("author_name")
	paperID := c.Request.FormValue("paper_id")
	paperTitle := c.Request.FormValue("paper_title")
	citation, _ := strconv.ParseUint(c.Request.FormValue("n_citation"), 0, 64)
	hindex, _ := strconv.ParseInt(c.Request.FormValue("hindex"), 0, 64)
	reason := c.Request.FormValue("reason")
	paperRecommend, notFound := service.QueryARecommendInPaperRecommend(paperID)
	if notFound {
		service.AddToPaperRecommend(paperID, paperTitle, citation, hindex)
	} else {
		service.UpdatePaperRecommend(&paperRecommend, hindex)
	}
	if err := service.CreateARecommend(authorID, authorName, paperID, citation, reason); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "推荐成功"})
	return
}

// RemoveRecommend doc
// @description 删除某条推荐
// @Tags portal
// @Param author_id formData string true "作者ID"
// @Param paper_id formData string true "文献ID"
// @Param hindex formData string true "h-index"
// @Success 200 {string} string "{"success": true, "message": "删除成功"}"
// @Router /portal/column/remove [post]
func RemoveRecommend(c *gin.Context) {
	authorID := c.Request.FormValue("author_id")
	paperID := c.Request.FormValue("paper_id")
	hindex, _ := strconv.ParseInt(c.Request.FormValue("hindex"), 0, 64)
	paperRecommend, _ := service.QueryARecommendInPaperRecommend(paperID)
	service.UpdatePaperRecommend(&paperRecommend, -hindex)
	if err := service.DeleteRecommend(authorID, paperID); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "删除成功"})
	return
}

// ListRecommendsFromOneAuthor doc
// @description 获取作者推荐的所有内容
// @Tags portal
// @Param author_id formData string true "作者ID"
// @Success 200 {string} string "{"success": true, "message": "查找成功", "data": "作者的所有推荐"}"
// @Router /portal/recommend/recommends_from_one_author [post]
func ListRecommendsFromOneAuthor(c *gin.Context) {
	authorID := c.Request.FormValue("author_id")
	recommends := service.QueryRecommendsFromOneAuthor(authorID)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查找成功", "data": recommends})
	return
}

// ListRecommendsFromOnePaper doc
// @description 获取所有对某文章的推荐
// @Tags portal
// @Param paper_id formData string true "文献ID"
// @Success 200 {string} string "{"success": true, "message": "查找成功", "data": "文献的所有推荐"}"
// @Router /portal/recommend/recommends_from_one_paper [post]
func ListRecommendsFromOnePaper(c *gin.Context) {
	paperID := c.Request.FormValue("paper_id")
	recommends := service.QueryRecommendsFromOnePaper(paperID)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查找成功", "data": recommends})
	return
}

// ListTopSevenPapers doc
// @description 获取推荐数最多的前七篇文献
// @Tags portal
// @Success 200 {string} string "{"success": true, "message": "查找成功", "data": "前七篇文献的ID"}"
// @Router /portal/recommend/top [post]
func ListTopSevenPapers(c *gin.Context) {
	papers := service.QueryTopSevenPapers()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查找成功", "data": papers})
	return
}

// IsSettled doc
// @description 判断该作者是否入驻
// @Tags portal
// @Param author_id formData string true "作者ID"
// @Success 200 {string} string "{"success": true, "message": "true"}"
// @Success 200 {string} string "{"success": true, "message": "false"}"
// @Router /portal/issettled [post]
func IsSettled(c *gin.Context) {
	authorID := c.Request.FormValue("author_id")
	_, notFound := service.FindPortalByID(authorID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "false"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "true"})
	}
	return
}
