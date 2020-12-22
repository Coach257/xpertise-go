package router

import (
	v1 "xpertise-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitPortalRouter(Router *gin.RouterGroup) {
	PortalRouter := Router.Group("api/v1/portal")
	{
		PortalRouter.POST("/author", v1.SearchAuthor)
		PortalRouter.POST("/is_settled", v1.IsSettled)
		PortalRouter.POST("/authorized_user_info", v1.AuthorizedUserInfo)

		PortalRouter.POST("/column/create_column", v1.CreateSpecialColumn)
		PortalRouter.POST("/column/add_to_column", v1.AddToColumn)
		PortalRouter.POST("/column/list_all_from_column", v1.ListAllFromAColumn)
		PortalRouter.POST("/column/remove_from_column", v1.RemovePaperFromColumn)
		PortalRouter.POST("/column/searchcol", v1.SearchSpecialColumn)

		PortalRouter.POST("/recommend/create", v1.CreateRecommend)
		PortalRouter.POST("/recommend/remove", v1.RemoveRecommend)
		PortalRouter.POST("/recommend/recommends_from_one_author", v1.ListRecommendsFromOneAuthor)
		PortalRouter.POST("/recommend/recommends_from_one_paper", v1.ListRecommendsFromOnePaper)
		PortalRouter.GET("/recommend/cs/top", v1.ListTopSevenCsPapers)
		PortalRouter.GET("/recommend/main/top", v1.ListTopSevenPapers)
		PortalRouter.POST("/direct_connection/list", v1.ListDirectConnectedAuthors)
		PortalRouter.POST("/author_connection_graph", v1.CreateAuthorConnectionsGraph)
	}
}
