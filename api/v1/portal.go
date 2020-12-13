package v1

import (
	"net/http"
	"xpertise-go/service"

	"github.com/gin-gonic/gin"
)

func QueryAPaperByID(c *gin.Context) {
	paper, err := service.QueryAPaperByID("000127EC")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": paper,
	})
	return

}

func TellAuthorInfo(c *gin.Context) {
	author, _ := service.QueryAnAuthorByID(c.Request.FormValue("author_id"))
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "查看作者信息成功",
		"data":    author,
	})
	return
}
