package controller

import (
	"strconv"
	"xpertise-go/user/dao"
	"xpertise-go/user/server"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Xpertise Scholar",
	})
}

func CreateAUser(c *gin.Context) {
	user := dao.User{Username: "Rolin", Password: "123456", Email: "1207640183@qq.com"}
	if err := server.CreateAUser(&user); err != nil {
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

func Register(c *gin.Context)  {
	/*
	request:
	{
		"username":string,
		"password":string,
		"password2":string,
		"email":string,
		"info":string
	}
	*/

	username :=c.Request.FormValue("username")
	password :=c.Request.FormValue("password")
	password2 :=c.Request.FormValue("password2")
	email :=c.Request.FormValue("email")
	info :=c.Request.FormValue("info")

	if server.QueryAUserByUsername(username) !=nil{
		c.JSON(200,gin.H{"success":false,"message":"用户名已被占用"})
	}
	if password != password2{
		c.JSON(200,gin.H{"success":false,"message":"两次密码不一致"})
	}
	if email == ""{
		c.JSON(200,gin.H{"success":false,"message":"未输入邮箱"})
	}
	if server.QueryAUserByEmail(email) !=nil {
		c.JSON(200,gin.H{"success":false,"message":"邮箱已被占用"})
	}

	user:=dao.User{Username: username,Password: password,Email:email,BasicInfo:info}
	server.CreateAUser(&user)
	c.JSON(200,gin.H{"success":true,"message":"用户创建成功"})
}
