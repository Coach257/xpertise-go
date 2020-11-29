package router

import (
	"xpertise-go/portal/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter contains all the api that will be used.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.Index)
	// api
	v2 := r.Group("api/portal")
	{
		v2.GET("/document/create", controller.CreateADocument)
		v2.GET("/document/query", controller.QueryDocumentByID)
	}
	return r
}
