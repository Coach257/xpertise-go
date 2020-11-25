package controller

import (
	"gin-project/dao"
	"gin-project/server"

	"github.com/gin-gonic/gin"
)

// Index func.
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Home Page.",
	})
}

// Test func.
func Test(c *gin.Context) {
	// 插入两条数据
	student1 := dao.Student{ID: 18373059, Name: "胡鹏飞", Age: 20}
	student2 := dao.Student{ID: 18373722, Name: "朱英豪", Age: 20}
	// 创建Table
	dao.CreateTableStudent()
	// 在表中插入数据
	server.CreateAStudent(&student1)
	server.CreateAStudent(&student2)

	// 删除数据
	//server.DeleteAStudentByID(18373059)

	// 这里要先定义再传student
	var student []dao.Student
	// 获取所有数据到 student中
	dao.DB.Find(&student)

	// 封装成 JSon文件
	c.IndentedJSON(200, student)
	//c.JSON(200, gin.H{"message": "create a user"})

	dao.DB.Model(&student).Where("Name = ?", "胡鹏飞").Update("Age", 21)
	//dao.DBModel(&student).Update("Age", 21)

	// 重新获取全部数据
	dao.DB.Find(&student)
	c.IndentedJSON(200, student)
}
