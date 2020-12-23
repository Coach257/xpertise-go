package v1

import (
	"net/http"
	"strconv"
	"xpertise-go/model"
	"xpertise-go/service"

	"github.com/gin-gonic/gin"
)

// Register doc
// @description 注册
// @Tags user
// @Param username formData string true "用户名"
// @Param password1 formData string true "密码1"
// @Param password2 formData string true "密码2"
// @Param email formData string true "邮箱"
// @Param info formData string true "个人信息"
// @Success 200 {string} string "{"success": true, "message": "用户创建成功"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	username := c.Request.FormValue("username")
	password1 := c.Request.FormValue("password1")
	password2 := c.Request.FormValue("password2")
	email := c.Request.FormValue("email")
	info := c.Request.FormValue("info")

	if password1 != password2 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "两次密码不一致",
		})
		return
	}

	_, notFoundUsername := service.QueryAUserByUsername(username)
	_, notFoundEmail := service.QueryAUserByEmail(email)
	if notFoundUsername && notFoundEmail {
		user := model.User{Username: username, Password: password1, Email: email, BasicInfo: info}
		service.CreateAUser(&user)
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "用户创建成功"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": false, "message": "用户名/邮箱已存在"})
	return
}

// Login doc
// @description 用户登录
// @Tags user
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"success": true, "message": "登录成功", "data": "model.User的所有信息"}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var user model.User
	var notFound bool

	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	user, notFound = service.QueryAUserByUsername(username)

	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户名不存在",
		})
		return
	}

	if user.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "密码错误",
		})
		return
	}

	//data, _ := jsoniter.Marshal(&user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "登录成功",
		"data":    user,
	})

	return
}

// ModifyUser doc
// @description 修改用户信息
// @Tags user
// @Param user_id formData string true "用户ID"
// @Param username formData string true "用户名"
// @Param password1 formData string true "原密码"
// @Param password2 formData string true "新密码"
// @Param email formData string true "邮箱"
// @Param info formData string true "个人信息"
// @Success 200 {string} string "{"success": true, "message": "登录成功", "data": "model.User的所有信息"}"
// @Router /user/modify [post]
func ModifyUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	username := c.Request.FormValue("username")
	password1 := c.Request.FormValue("password1")
	password2 := c.Request.FormValue("password2")
	email := c.Request.FormValue("email")
	info := c.Request.FormValue("info")
	user, notFoundUserByID := service.QueryAUserByID(userID)
	if notFoundUserByID {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户ID不存在",
		})
		return
	}
	if password1 != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "原密码输入错误",
		})
		return
	}
	_, notFoundUserByName := service.QueryAUserByUsername(username)
	if !notFoundUserByName && username != user.Username {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户名已被占用",
		})
		return
	}
	_, notFoundUserByEmail := service.QueryAUserByEmail(email)
	if !notFoundUserByEmail && email != user.Email {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "邮箱已被占用",
		})
		return
	}
	err := service.UpdateAUser(&user, username, password2, email, info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	//data, _ := jsoniter.Marshal(&user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "用户信息修改成功",
		"data":    user,
	})
	return
}

// TellUserInfo doc
// @description 查看用户个人信息
// @Tags user
// @Param user_id formData string true "用户ID"
// @Success 200 {string} string "{"success": true, "message": "查看用户信息成功", "data": "model.User的所有信息"}"
// @Router /user/info [post]
func TellUserInfo(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user, _ := service.QueryAUserByID(userID)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "查看用户信息成功",
		"data":    user,
	})
	return
}

// AddToFavorites doc
// @description 添加收藏
// @Tags user
// @Param user_id formData string true "用户ID"
// @Param paper_id formData string true "文献ID"
// @Param paper_info formData string true "文献描述"
// @Success 200 {string} string "{"success":true, "message":"收藏成功"}"
// @Router /user/favorite/add [post]
func AddToFavorites(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	paperID := c.Request.FormValue("paper_id")
	paperInfo := c.Request.FormValue("paper_info")

	if userID==0{
		c.JSON(http.StatusOK,gin.H{"success":false,"message":"请先登录"})
		return
	}

	if err := service.CreateAFavorite(userID, paperID, paperInfo); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "收藏成功"})
	}
}

// ListAllFavorites doc
// @description 获取收藏列表
// @Tags user
// @Param user_id formData string true "用户ID"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"user的所有收藏"}"
// @Router /user/favorite/list [post]
func ListAllFavorites(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user := service.QueryAllFavorites(userID)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": user.Favorites})
}

// RemoveFavorite doc
// @description 移除收藏
// @Tags user
// @Param favor_id formData string true "收藏ID"
// @Success 200 {string} string "{"success":true, "message":"删除成功"}"
// @Router /user/favorite/remove [post]
func RemoveFavorite(c *gin.Context) {
	favorID, _ := strconv.ParseUint(c.Request.FormValue("favor_id"), 0, 64)
	if err := service.DeleteAFavorite(favorID); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "删除成功"})
	}
	return
}

// IsFavorite doc
// @description 判断是否已经被收藏
// @Tags user
// @Param user_id formData string true "用户ID"
// @Param paper_id formData string true "文献ID"
// @Success 200 {string} string "{"success":true, "message":"true"}"
// @Router /user/favorite/isfav [post]
func IsFavorite(c *gin.Context) {
	userID := c.Request.FormValue("user_id")
	paperID := c.Request.FormValue("paper_id")
	_, notFound := service.FindFavByUPID(userID, paperID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "false"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "true"})
	}
}

// IsUserWish doc
// @description 判断该篇Paper是否已在用户心愿清单中
// @Tags user
// @Param user_id formData string true "用户ID"
// @Param paper_id formData string true "文献ID"
// @Success 200 {string} string "{"success":true, "message":"不在用户的心愿清单内/已在用户的心愿清单中"}"
// @Router /user/wish/paper_in_wish [post]
func IsUserWish(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	paperID := c.Request.FormValue("paper_id")
	_, notFound := service.QueryAWish(userID, paperID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "不在用户的心愿清单内"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "已在用户的心愿清单中"})
	return
}

// AddToWishes doc
// @description 添加至心愿清单
// @Tags user
// @Param user_id formData string true "用户ID"
// @Param paper_id formData string true "文献ID"
// @Param title formData string true "文献标题"
// @Param year formData string true "PaperPublishYear"
// @Param n_citation formData string true "Paper引用数量"
// @Success 200 {string} string "{"success":true, "message":"已添加至心愿清单"}"
// @Router /user/wish/add [post]
func AddToWishes(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	paperID := c.Request.FormValue("paper_id")
	title := c.Request.FormValue("title")
	year := c.Request.FormValue("year")
	citation, _ := strconv.ParseUint(c.Request.FormValue("n_citation"), 0, 64)
	url := c.Request.FormValue("url")
	if err := service.CreateAWish(userID, paperID, title, url, year, citation); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "已添加至心愿清单"})
	return
}

// RemoveFromWishes doc
// @description 移出心愿清单
// @Tags user
// @Param wish_id formData string true "心愿ID"
// @Success 200 {string} string "{"success":true, "message":"已移出清单"}"
// @Router /user/wish/remove [post]
func RemoveFromWishes(c *gin.Context) {
	wishID, _ := strconv.ParseUint(c.Request.FormValue("wish_id"), 0, 64)
	if err := service.DeleteAWish(wishID); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "已移出心愿清单"})
}

// ListAllWishes doc
// @description 获取心愿清单列表
// @Tags user
// @Param user_id formData string true "用户ID"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"user的清单"}"
// @Router /user/wish/list [post]
func ListAllWishes(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user := service.QueryAllWishes(userID)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": user.Wishes})

}

func DeleteAUserByID(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	err := service.DeleteAUserByID(userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "删除用户失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "删除用户成功",
	})
}

// GetUserAllAuthorizationRequest doc
// @description 获取用户的（所有）请求认证
// @Tags user
// @Param user_id formData string true "用户ID"
// @Success 200 {string} string "{"success": true, "message": "获取申请信息成功。", "data": "请求认证的所有信息。"}"
// @Router /user/authorize/get [post]
func GetUserAllAuthorizationRequest(c *gin.Context) {
	userIDStr := c.Request.FormValue("user_id")
	userID, _ := strconv.ParseUint(userIDStr, 0, 64)
	aureqs, notFound := service.QueryAuthorizationRequestsByUserID(userID)
	if notFound == true {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "暂无申请信息。",
			"data":    aureqs,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "获取申请信息成功。",
			"data":    aureqs,
		})
	}
	return
}

// ReadAUserAuthorizationRequest doc
// @description 已读一条请求认证
// @Tags user
// @Param user_id formData string true "用户ID"
// @Param authreq_id formData string true "请求认证ID"
// @Success 200 {string} string "{"success": true, "message": "标记已读成功！"}"
// @Router /user/authorize/read [post]
func ReadAUserAuthorizationRequest(c *gin.Context) {
	userIDStr := c.Request.FormValue("user_id")
	userID, _ := strconv.ParseUint(userIDStr, 0, 64)
	authreqIDStr := c.Request.FormValue("authreq_id")
	authreqID, _ := strconv.ParseUint(authreqIDStr, 0, 64)
	err := service.DeleteAnAuthorizationRequest(authreqID, userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "标记已读失败。",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "标记已读成功！",
		})
	}
}
