package v1

import (
	"net/http"
	"strconv"
	"xpertise-go/model"
	"xpertise-go/service"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
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
			"message": "用户不存在",
		})
		return
	}

	if user.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户名或密码错误",
		})
		return
	}

	data, _ := jsoniter.Marshal(&user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "登录成功",
		"data":    string(data),
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
// @Router /user/modify_user [post]
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
			"message": "用户不存在",
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
	data, _ := jsoniter.Marshal(&user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "用户信息修改成功",
		"data":    string(data),
	})
	return
}
