package router

import (
	"net/http"
	adminController "xpertise-go/admin/controller"
	branchController "xpertise-go/branch/controller"
	portalController "xpertise-go/portal/controller"
	"xpertise-go/user/auth"
	userController "xpertise-go/user/controller"

	"github.com/gin-gonic/gin"
)

//Cors solve cors problem.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// SetupRouter contains all the api that will be used.
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	adminV1 := r.Group("api/v1/admin")
	{
		adminV1.POST("/forbid/user", adminController.ForbidAUser)
		adminV1.POST("/report/comment", adminController.DealWithAComReport)
	}

	branchV1 := r.Group("api/v1/branch")
	{
		branchV1.POST("/create", branchController.CreateAComment)             //post requires comid userid content
		branchV1.DELETE("/delete/id", branchController.DeleteACommentByID)    //delete/3
		branchV1.POST("/thumbup", branchController.AddLike)                   //post requires docid userid  两个参数
		branchV1.POST("/revert_thumbup", branchController.RevertAddLike)      // post requires docid userid  两个参数 确保已经点赞
		branchV1.POST("/thumbdown", branchController.AddDisLike)              // post requires docid userid  两个参数
		branchV1.POST("/revert_thumbdown", branchController.RevertAddDisLike) // post requires docid userid  两个参数
		branchV1.GET("/thumbup/query", branchController.QueryThumbUp)         //  requires id  评论id参数
		branchV1.GET("/thumbdown/query", branchController.QueryThumbDown)     // requires id  评论id参数

	}

	portalV1 := r.Group("api/v1/portal")
	{
		portalV1.POST("/doc/create", portalController.CreateDocument)
		portalV1.GET("/doc/query/id", portalController.QueryDocumentByID)
		portalV1.GET("/doc/query/title", portalController.QueryDocumentsByTitle)
		portalV1.POST("/org/create", portalController.CreateOrganization)
		portalV1.GET("/org/query/id", portalController.QueryOrganizationByID)
		portalV1.GET("/org/query/name", portalController.QueryOrganizationByName)
	}

	userV1 := r.Group("api/v1/user")
	{
		userV1.DELETE("/delete/id", userController.DeleteAStudentByID)
		userV1.PUT("/update/age", userController.UpdateAStudentByAge)
		userV1.GET("/query/all", userController.QueryAllUsers)
		userV1.GET("/query/id", userController.QueryStudentByID)
		userV1.GET("/query/age", userController.QueryStudentsByAge)
		userV1.POST("/register", userController.Register)
		userV1.POST("/login", userController.Login)
		userV1.POST("/reset/password", userController.ResetPassword)
		userV1.GET("/return/account_info",userController.ReturnAccountInfo)

		userV1.POST("/reset/account_info", auth.JwtAuth(), userController.ResetAccountInfo)
		userV1.POST("/folder/create", auth.JwtAuth(), userController.CreateAFolder)
		userV1.POST("/folder/add", auth.JwtAuth(), userController.AddToMyFolder)

	}
	return r
}
