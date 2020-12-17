package router

import (
	v1 "xpertise-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitPortalRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("api/v1/portal")
	{
		UserRouter.POST("/create_column", v1.CreateSpecialColumn)
		UserRouter.POST("/add_to_column", v1.AddToColumn)
		UserRouter.POST("/list_all_from_column", v1.ListAllFromAColumn)
		UserRouter.POST("/remove_from_column", v1.RemovePaperFromColumn)
		UserRouter.POST("/author", v1.SearchAuthor)
	}
}
