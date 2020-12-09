package router

import (
	v1 "xpertise-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("api/v1/user")
	{
		UserRouter.POST("/register", v1.Register)
		UserRouter.POST("/login", v1.Login)
		UserRouter.POST("/modify", v1.ModifyUser)
		UserRouter.POST("/create_folder", v1.CreateAFolder)
		UserRouter.POST("/info", v1.TellUserInfo)
		UserRouter.POST("delete", v1.DeleteAUserByID)
	}
}
