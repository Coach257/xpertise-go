package controller

import (
	"time"
	"xpertise-go/dao"
	"xpertise-go/portal/server"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Xpertise Scholar",
	})
}

func CreateDocument(c *gin.Context) {
	var doc dao.Document
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(0, gin.H{"message": err})
	}
	doc.PublishTime = time.Now()
	if err := server.DocCreate(&doc); err != nil {
		c.JSON(0, gin.H{"message": err})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}
func CreateOrganization(c *gin.Context) {
	var org dao.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(0, gin.H{"message": err})
	}
	if err := server.OrgCreate(&org); err != nil {
		c.JSON(0, gin.H{"message": err})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}

type queryParams struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

func QueryDocumentByID(c *gin.Context) {
	//Form-data version
	//arg, _ := strconv.ParseUint(c.PostForm("id"), 0, 64)
	//document := server.DocQueryID(arg)

	//json version
	var arg queryParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(0, gin.H{"message": err})
	}
	document := server.DocQueryID(arg.Id)
	c.IndentedJSON(200, document)
}

func QueryOrganizationByID(c *gin.Context) {
	//Form-data version
	//arg, _ := strconv.ParseUint(c.PostForm("id"), 0, 64)
	//organization := server.OrgQueryID(arg)

	//json version
	var arg queryParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(0, gin.H{"message": err})
	}
	organization := server.OrgQueryID(arg.Id)

	c.IndentedJSON(200, organization)
}

func QueryDocumentsByTitle(c *gin.Context) {
	var arg queryParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(0, gin.H{"message": err})
	}
	documents := server.DocQueryTitle(arg.Title)

	c.IndentedJSON(200, documents)
}

func QueryOrganizationByName(c *gin.Context) {
	var arg queryParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(0, gin.H{"message": err})
	}
	organizations := server.OrgQueryName(arg.Name)

	c.IndentedJSON(200, organizations)
}
