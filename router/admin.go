package router

import (
	v1 "xpertise-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitAdminRouter(Router *gin.RouterGroup) {
	AdminRouter := Router.Group("api/v1/admin")
	{
		AdminRouter.POST("/authorize/request", v1.RequestForAuthorization)
		AdminRouter.POST("/authorize/deal", v1.DealWithAuthorizationRequest)
		AdminRouter.GET("/authorize/all", v1.GetAllAuthorizationRequest)
	}
}
