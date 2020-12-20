package router

import (
	v1 "xpertise-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBranchRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("api/v1/branch")
	{
		UserRouter.POST("/comment/create", v1.CreateAComment) // 注意不用加逗号
		UserRouter.POST("/comment/operate", v1.OperateComment)
		UserRouter.POST("/comment/give_a_like_or_dislike", v1.GiveALikeOrDislike)
		UserRouter.POST("/graph/author_connection", v1.AuthorConnection)
		UserRouter.POST("/comment/list_all_comments", v1.ListAllComments)
		UserRouter.POST("/graph/reference", v1.GetThreeLevelReferences)
	}
}
