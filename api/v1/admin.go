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
// @Param citizen_id formData string true "身份证号"
// @Param organization formData string true "工作单位"
// @Success 200 {string} string "{"success": true, "message": "申请认证成功。"}"
// @Router /admin/authorize/request [post]
func RequestForAuthorization(c *gin.Context) {
	userIDStr := c.Request.FormValue("user_id")
	userID, _ := strconv.ParseUint(userIDStr, 0, 64)
	citizenID := c.Request.FormValue("citizen_id")
	organization := c.Request.FormValue("organization")
	if service.CreateAnAuthorizationRequest(userID, citizenID, organization) != nil {
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
// @Param author_id formData string true "对应作者ID"
// @Success 200 {string} string "{"success": true, "message": "已通过认证请求。"}"
// @Router /admin/authorize/deal [post]
func DealWithAuthorizationRequest(c *gin.Context) {
	authreqIDStr := c.Request.FormValue("authreq_id")
	if authreqIDStr == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "authreq_id为空。",
		})
		return
	}
	authreqID, _ := strconv.ParseUint(authreqIDStr, 0, 64)
	action := c.Request.FormValue("action")
	if action == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "action为空。",
		})
		return
	}
	authreq, notFound := service.QueryAnAuthorizationRequest(authreqID)
	if notFound == true {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "没有对应的请求。",
		})
		return
	}
	if action == "Accept" {
		authorID := c.Request.FormValue("author_id")
		service.UpdateAnAuthorizationRequest(&authreq, "Accepted", authorID)
		service.CreateAPortal(authreq.UserID, authreq.AuthorID)
		userID := authreq.UserID
		user, notFound := service.QueryAUserByID(userID)
		if notFound {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "用户不存在。",
			})
		} else {
			// 更新UserType
			service.UpdateAUserUserType(&user, 2)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "已通过认证请求。",
			})
		}
	} else if action == "Reject" {
		service.UpdateAnAuthorizationRequest(&authreq, "Rejected", "")
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "已拒绝认证请求。",
		})
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
