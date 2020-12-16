package v1

import (
	"net/http"
	"strconv"
	"xpertise-go/service"

	"github.com/gin-gonic/gin"
)

// RequestForAuthorization doc
// @description 发送请求认证
// @Tags admin
// @Param user_id formData string true "用户名"
// @Param author_id formData string true "作者ID"
// @Success 200 {string} string "{"success": true, "message": "申请认证成功。"}"
// @Router /admin/authorize/request [post]
func RequestForAuthorization(c *gin.Context) {
	userIDStr := c.Request.FormValue("user_id")
	userID, _ := strconv.ParseUint(userIDStr, 0, 64)
	authorID := c.Request.FormValue("author_id")
	if service.CreateAuthorizationRequest(userID, authorID) == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "申请认证失败。",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "申请认证成功。",
		})
	}
	return
}

// DealWithAuthorizationRequest doc
// @description 处理用户申请认证
// @Tags admin
// @Param authreq_id formData string true "用户申请认证ID"
// @Param action formData string true "Accept/Reject"
// @Success 200 {string} string "{"success": true, "message": "已通过认证请求。"}"
// @Router /admin/authorize/deal [post]
func DealWithAuthorizationRequest(c *gin.Context) {
	authreqIDStr := c.Request.FormValue("authreq_id")
	authreqID, _ := strconv.ParseUint(authreqIDStr, 0, 64)
	action := c.Request.FormValue("action")
	authreq, err := service.QueryAAuthorizationRequest(authreqID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "没有对应的请求。",
		})
	}
	if action == "Accept" {
		service.CreateAPortal(authreq.UserID, authreq.AuthorID)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "已通过认证请求。",
		})
		service.DeleteAuthorizationRequest(authreqID)
	} else if action == "Reject" {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "已拒绝认证请求。",
		})
		service.DeleteAuthorizationRequest(authreqID)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Action字段无效",
		})
	}
	return
}

// GetAllAuthorizationRequest doc
// @description 获取用户申请条目
// @Tags admin
// @Success 200 {string} string "{"success": true, "message": "获取条目成功。", "data": "model.AuthorizationRequest的所有信息"}"
// @Router /admin/authorize/all [get]
func GetAllAuthorizationRequest(c *gin.Context) {
	authreqs := service.QueryAllAuthorizationRequest()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取条目成功。",
		"data":    authreqs,
	})
	return
}
