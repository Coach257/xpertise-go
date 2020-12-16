package initialize

import (
	"xpertise-go/middleware"
	"xpertise-go/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	Group := r.Group("")
	{
		router.InitUserRouter(Group) // 注册用户路由
		router.InitPortalRouter(Group)
		router.InitBranchRouter(Group)
		router.InitAdminRouter(Group)
	}
	return r
}
