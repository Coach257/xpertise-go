package controller

import (
	"xpertise-go/branch/server"
	"xpertise-go/dao"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Xpertise Scholar",
	})

}

func CreateAComment(c *gin.Context) {
	com := dao.Comment{UserID: 123, DocID: 11, Content: "这好吗"}
	if err := server.CreateAComment(&com); err != nil {
		c.JSON(0, gin.H{"message": err})
	} else {
		c.JSON(200, gin.H{"message": "CommentAddSuccess"})
	}
}
