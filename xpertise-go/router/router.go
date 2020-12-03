package router

import (
	branchController "xpertise-go/branch/controller"
	portalController "xpertise-go/portal/controller"
	"xpertise-go/user/auth"
	userController "xpertise-go/user/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter contains all the api that will be used.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// adminV1 := r.Group("api/v1/admin")

	branchV1 := r.Group("api/v1/branch")
	{
		branchV1.POST("/create", branchController.CreateAComment)             //post requires comid userid content
		branchV1.DELETE("/delete/:id", branchController.DeleteACommentByID)   //delete/3
		branchV1.POST("/thumbup", branchController.AddLike)                   //post requires docid userid  两个参数
		branchV1.POST("/revert_thumbup", branchController.RevertAddLike)      // post requires docid userid  两个参数 确保已经点赞
		branchV1.POST("/thumbdown", branchController.AddDisLike)              // post requires docid userid  两个参数
		branchV1.POST("/revert_thumbdown", branchController.RevertAddDisLike) // post requires docid userid  两个参数
		branchV1.GET("/query/thumbup", branchController.QueryThumbUp)         //  requires id  评论id参数
		branchV1.GET("/query/thumbdown", branchController.QueryThumbDown)     // requires id  评论id参数

	}

	portalV1 := r.Group("api/v1/portal")
	{
		portalV1.POST("/document/create", portalController.CreateADocument)
		portalV1.GET("/document/query", portalController.QueryDocumentByID)
	}
	userV1 := r.Group("api/v1/user")
	{
		userV1.DELETE("/delete/:id", userController.DeleteAStudentByID)
		userV1.PUT("/update/age", userController.UpdateAStudentByAge)
		userV1.GET("/query/all", userController.QueryAllStudents)
		userV1.GET("/query/id", userController.QueryStudentByID)
		userV1.GET("/query/age", userController.QueryStudentsByAge)
		userV1.POST("/register", userController.Register)
		userV1.POST("/login", userController.Login)

		userV1.POST("/password/reset", userController.ResetPassword)
		userV1.POST("/folder/create", auth.JwtAuth(), userController.CreateAFolder)
		userV1.POST("/folder/add", auth.JwtAuth(), userController.AddToMyFolder)
		userV1.POST("/account_info/reset", userController.ResetAccountInfo)
	}
	return r
}
