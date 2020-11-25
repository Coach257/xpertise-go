package controller

import (
	"gin-project/dao"
	"gin-project/server"
	"time"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	user := dao.User{ID: time.Now().Unix(), Name: "施哲纶"}
	server.CreateAUser(&user)
	c.JSON(200, gin.H{
		"message": "create a user",
	})
}
