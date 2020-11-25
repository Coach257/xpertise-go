package router

import (
	"gin-project/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter contains all the api that will be used.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/index", controller.Index)
	// api
	apiGroup := r.Group("api")
	{
		apiGroup.GET("/test", controller.Test)
	}
	return r
}
