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
	test := r.Group("api/portal/test")
	{
		test.GET("/document/create", controller.CreateADocument)
		test.GET("/document/query", controller.QueryDocumentByID)
	}

	return r
}
