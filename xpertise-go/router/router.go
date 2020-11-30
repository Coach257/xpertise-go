package router

import (
	adminController "xpertise-go/admin/controller"
	branchController "xpertise-go/branch/controller"
	portalController "xpertise-go/portal/controller"
	userController "xpertise-go/user/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter contains all the api that will be used.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	adminV1 := r.Group("api/v1/admin")
	{
		adminV1.POST("/create", adminController.CreateAStudent)
		adminV1.DELETE("/delete/:id", adminController.DeleteAStudentByID)
		adminV1.PUT("/update/age", adminController.UpdateAStudentByAge)
		adminV1.GET("/query/all", adminController.QueryAllStudents)
		adminV1.GET("/query/id", adminController.QueryStudentByID)
		adminV1.GET("/query/age", adminController.QueryStudentsByAge)
	}
	branchV1 := r.Group("api/v1/branch")
	{
		branchV1.POST("/create", branchController.CreateAStudent)
		branchV1.DELETE("/delete/:id", branchController.DeleteAStudentByID)
		branchV1.PUT("/update/age", branchController.UpdateAStudentByAge)
		branchV1.GET("/query/all", branchController.QueryAllStudents)
		branchV1.GET("/query/id", branchController.QueryStudentByID)
		branchV1.GET("/query/age", branchController.QueryStudentsByAge)
	}
	portalV1 := r.Group("api/v1/portal")
	{
		portalV1.POST("/document/create", portalController.CreateADocument)
		portalV1.GET("/document/query", portalController.QueryDocumentByID)
	}
	userV1 := r.Group("api/v1/user")
	{
		userV1.POST("/create", userController.CreateAUser)
		userV1.DELETE("/delete/:id", userController.DeleteAStudentByID)
		userV1.PUT("/update/age", userController.UpdateAStudentByAge)
		userV1.GET("/query/all", userController.QueryAllStudents)
		userV1.GET("/query/id", userController.QueryStudentByID)
		userV1.GET("/query/age", userController.QueryStudentsByAge)
		userV1.POST("/register",userController.Register)
	}

	return r
}
