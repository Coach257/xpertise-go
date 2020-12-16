package router

import (
	v1 "xpertise-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitPortalRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("api/v1/portal")
	{
		UserRouter.POST("/create_column", v1.CreateSpecialColumn)
	}
}
