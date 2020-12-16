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
		UserRouter.POST("/info", v1.TellUserInfo)
		UserRouter.POST("/delete", v1.DeleteAUserByID)
		UserRouter.POST("/favorite/add", v1.AddToFavorites)
		UserRouter.POST("/favorite/list", v1.ListAllFavorites)
		UserRouter.POST("/favorite/remove", v1.RemoveFavorite)
		UserRouter.POST("/wish/add", v1.AddToWishes)
		UserRouter.POST("/wish/remove", v1.RemoveFromWishes)
		UserRouter.POST("/wish/list", v1.ListAllWishes)
	}
}
