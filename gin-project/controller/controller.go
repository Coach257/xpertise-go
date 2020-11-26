package controller

import (
	"gin-project/dao"
	"gin-project/server"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Xpertise Scholar",
	})
}

func CreateAStudent(c *gin.Context) {
	stu := dao.Student{ID: 1, Name: "hpf", Age: 20}
	if err := server.CreateAStudent(&stu); err != nil {
		c.JSON(0, gin.H{"message": err})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}

func DeleteAStudentByID(c *gin.Context) {
	sid, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	server.DeleteAStudentByID(sid)
	c.JSON(200, gin.H{"message": "success"})
}

func QueryAllStudents(c *gin.Context) {
	students := server.QueryAllStudents()
	c.IndentedJSON(200, students)
}

func UpdateAStudentByAge(c *gin.Context) {
	sid, _ := strconv.ParseUint(c.PostForm("sid"), 0, 64)
	age, _ := strconv.ParseUint(c.PostForm("age"), 0, 64)
	student := server.QueryStudentByID(sid)
	server.UpdateAStudentByAge(student[0], age)
	c.JSON(200, gin.H{"message": "success"})
}

func QueryStudentByID(c *gin.Context) {
	sid, _ := strconv.ParseUint(c.PostForm("id"), 0, 64)
	student := server.QueryStudentByID(sid)
	c.IndentedJSON(200, student)
}

func QueryStudentsByAge(c *gin.Context) {
	age, _ := strconv.ParseUint(c.PostForm("age"), 0, 64)
	student := server.QueryStudentsByAge(age)
	c.IndentedJSON(200, student)
}
