package controller

import (
	"net/http"
	"strconv"
	"time"
	AdminServer "xpertise-go/admin/server"
	UserServer "xpertise-go/user/server"

	"github.com/gin-gonic/gin"
)

// DealWithAComReport deal with comment report
// Accept: delete
// Reject: ignore
func DealWithAComReport(c *gin.Context) {
	reportidStr := c.Request.FormValue("ReportID")
	reportid, _ := strconv.ParseUint(reportidStr, 0, 64)
	action := c.Request.FormValue("Action")
	comReport, notfound := AdminServer.QueryAComReportByID(reportid)
	if notfound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "举报条目不存在",
		})
		return
	}
	if action == "Accept" {
		AdminServer.UpdateAComReportByStatus(&comReport, action)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "已删除被举报评论",
		})
	} else if action == "Reject" {
		AdminServer.UpdateAComReportByStatus(&comReport, action)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "已驳回举报",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Action字段无效",
		})
	}

	return
}

// ForbidAUser 禁言用户
func ForbidAUser(c *gin.Context) {
	useridStr := c.Request.FormValue("UserID")
	userid, _ := strconv.ParseUint(useridStr, 0, 64)
	durationStr := c.Request.FormValue("Duration")
	duration, _ := time.ParseDuration(durationStr)
	user, notfound := UserServer.QueryAUerById(userid)
	if notfound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户不存在",
		})
		return
	}
	AdminServer.ForbidAUser(&user, &duration)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "禁言用户成功"})
	return
}
