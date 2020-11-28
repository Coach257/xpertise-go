package router

import (
	"xpertise-go/admin/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter contains all the api that will be used.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.Index)
	// api
	v1 := r.Group("api/v1/student")
	{
		v1.POST("/create", controller.CreateAStudent)
		v1.DELETE("/delete/:id", controller.DeleteAStudentByID)
		v1.PUT("/update/age", controller.UpdateAStudentByAge)
		v1.GET("/query/all", controller.QueryAllStudents)
		v1.GET("/query/id", controller.QueryStudentByID)
		v1.GET("/query/age", controller.QueryStudentsByAge)
	}
	return r
}
