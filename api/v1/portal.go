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
// @Router /portal/create_column [post]
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

// AddToColumn doc
// @description 添加某篇文章到专栏
// @Tags portal
// @Param paper_id formData string true "文献ID"
// @Param column_id formData string true "专栏ID"
// @Success 200 {string} string "{"success":true, "message":"添加到专栏成功"}"
// @Router /portal/add_to_column [post]
func AddToColumn(c *gin.Context) {
	columnID, _ := strconv.ParseUint(c.Request.FormValue("column_id"), 0, 64)
	paperID := c.Request.FormValue("paper_id")
	_, notFound := service.QueryItemFromColumnPaper(columnID, paperID)

	if !notFound {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "不能重复添加"})
		return
	}

	if err := service.AddPaperToColumn(columnID, paperID); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "添加到专栏成功"})
	return
}
