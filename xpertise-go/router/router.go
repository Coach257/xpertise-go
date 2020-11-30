package router

import (
	portalController "xpertise-go/portal/controller"
	userController "xpertise-go/user/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter contains all the api that will be used.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// adminV1 := r.Group("api/v1/admin")

	// branchV1 := r.Group("api/v1/branch")

	portalV1 := r.Group("api/v1/portal")
	{
		portalV1.GET("/document/query", portalController.QueryDocumentByID)
	}
	userV1 := r.Group("api/v1/user")
	{
		userV1.POST("/create", userController.CreateAUser)
		// userV1.DELETE("/delete/:id", userController.DeleteAStudentByID)
		// userV1.PUT("/update/age", userController.UpdateAStudentByAge)
		// userV1.GET("/query/all", userController.QueryAllStudents)
		// userV1.GET("/query/id", userController.QueryStudentByID)
		// userV1.GET("/query/age", userController.QueryStudentsByAge)
	}

	return r
}
