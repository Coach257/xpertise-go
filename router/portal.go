package router

import (
	v1 "xpertise-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitPortalRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("api/v1/portal")
	{
		UserRouter.POST("/query_paper", v1.QueryAPaperByID)
		UserRouter.POST("/author_info", v1.TellAuthorInfo)
	}
}
