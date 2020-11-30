package controller

import (
	"time"
	"xpertise-go/portal/dao"
	"xpertise-go/portal/server"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Xpertise Scholar",
	})
}

func CreateADocument(c *gin.Context) {
	doc := dao.Document{
		DocID:       1,
		AuthorList:  []string{"wyh", "yp"},
		TypeList:    []string{},
		Abstract:    "ahahha",
		CiteList:    []string{"cite1", "cite2"},
		PublishTime: time.Now(),
		Source:      "test",
		Original:    "abcjfdksleinxsijfiengklsjakl",
	}

	if err := server.CreateADocument(&doc); err != nil {
		c.JSON(0, gin.H{"message": err})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}

}

func QueryDocumentByID(c *gin.Context) {
	//id, _ := strconv.ParseUint(c.PostForm("id"), 0, 64)
	student := server.QueryDocument(1)
	c.IndentedJSON(200, student)
}
