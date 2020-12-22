package router

import (
	v1 "xpertise-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBranchRouter(Router *gin.RouterGroup) {
	BranchRouter := Router.Group("api/v1/branch")
	{
		BranchRouter.POST("/comment/create", v1.CreateAComment) // 注意不用加逗号
		BranchRouter.POST("/comment/operate", v1.OperateComment)
		BranchRouter.POST("/comment/give_a_like_or_dislike", v1.GiveALikeOrDislike)
		//BranchRouter.POST("/graph/author_connection", v1.AuthorConnection)
		BranchRouter.POST("/comment/list_all_comments", v1.ListAllComments)
		BranchRouter.POST("/graph/reference", v1.GetThreeLevelReferences)
		BranchRouter.POST("/reference_connection_graph", v1.ReferenceConnectionGraph)
	}
}
