package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"xpertise-go/dao"
	"xpertise-go/user/auth"
	"xpertise-go/user/server"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Xpertise Scholar",
	})
}

func DeleteAStudentByID(c *gin.Context) {
	sid, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	server.DeleteAStudentByID(sid)
	c.JSON(200, gin.H{"message": "success"})
}

func QueryAllUsers(c *gin.Context) {
	users := server.QueryAllUsers()
	c.IndentedJSON(200, users)
}

func UpdateAStudentByAge(c *gin.Context) {
	sid, _ := strconv.ParseUint(c.PostForm("sid"), 0, 64)
	age, _ := strconv.ParseUint(c.PostForm("age"), 0, 64)
	student := server.QueryStudentByID(sid)
	server.UpdateAStudentByAge(student[0], age)
	c.JSON(200, gin.H{"message": "success"})
}

func QueryStudentByID(c *gin.Context) {
	sid, _ := strconv.ParseUint(c.PostForm("id"), 0, 64)
	student := server.QueryStudentByID(sid)
	c.IndentedJSON(200, student)
}

func QueryStudentsByAge(c *gin.Context) {
	age, _ := strconv.ParseUint(c.PostForm("age"), 0, 64)
	student := server.QueryStudentsByAge(age)
	c.IndentedJSON(200, student)
}

func GenerateToken(c *gin.Context, user dao.User) (string, error) {
	j := &auth.JWT{
		[]byte("buaa21xpertise"),
	}
	claims := auth.CustomClaims{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), //签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), //过期时间 一小时
			Issuer:    "buaa21xpertise",                //签名发行者
		},
	}

	//创建一个token
	token, err := j.CreateToken(claims)

	return token, err
}

func InputBlankCheck(username string, email string) (bool, string) {
	if username == "" {
		return true, "未输入用户名"
	}
	if email == "" {
		return true, "未输入邮箱"
	}
	return false, ""
}

func DuplicateCheck(c *gin.Context, userId uint64, username string, email string, funcName string) bool {
	userFoundByUsername, notFoundName := server.QueryAUserByUsername(username)
	userFoundByEmail, notFoundEmail := server.QueryAUserByEmail(email)

	if funcName == "Register" {
		if !notFoundName {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "用户名已被占用"})
			return true
		}
		if !notFoundEmail {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "邮箱已被占用"})
			return true
		}
	}

	if funcName == "ResetAccountInfo" {
		if userFoundByUsername.UserID != userId {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "用户名已被占用"})
			return true
		}
		if userFoundByEmail.UserID != userId {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "邮箱已被占用"})
			return true
		}
	}
	return false
}

//用户注册
/*
	request:
	{
		"username":string,
		"password":string,
		"password2":string,
		"email":string,
		"info":string
	}
*/
func Register(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	password2 := c.Request.FormValue("password2")
	email := c.Request.FormValue("email")
	info := c.Request.FormValue("info")

	if blank, message := InputBlankCheck(username, email); blank {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": message,
		})
	}
	if duplicate := DuplicateCheck(c, 0, username, email, "Register"); duplicate {
		return
	}
	if password != password2 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "两次密码不一致",
		})
		return
	}

	user := dao.User{Username: username, Password: password, Email: email, BasicInfo: info}
	server.CreateAUser(&user)
	c.JSON(200, gin.H{"success": true, "message": "用户创建成功"})
}

//验证账户信息
func AccountCheck(c *gin.Context) (bool, dao.User) {
	var user dao.User
	var notfound bool

	username := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	//debug
	fmt.Println(username)
	fmt.Println(email)
	fmt.Println(password)

	if username != "" {
		user, notfound = server.QueryAUserByUsername(username)
	} else if email != "" {
		user, notfound = server.QueryAUserByEmail(email)
	} else {
		c.JSON(200, gin.H{
			"success": false,
			"message": "请输入邮箱或用户名",
		})
		return false, user
	}

	if notfound {
		c.JSON(200, gin.H{
			"success": false,
			"message": "用户或邮箱不存在",
		})
		return false, user
	}

	if user.Password != password {
		c.JSON(200, gin.H{
			"success": false,
			"message": "用户名或密码错误",
		})
		return false, user
	}

	return true, user
}

type LoginResult struct {
	Token        string `json:"token"`
	UserId       uint64 `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	UserType     int    `json:"user_type"`
	Info         string `json:"info"`
	Interdiction bool   `json:"interdiction"`
}

//用户登录
/*
	request:
	{
		"username":string,
		"email":string,
		"password":string
	}
*/
func Login(c *gin.Context) {
	pass, user := AccountCheck(c)

	if !pass {
		return
	}

	token, err := GenerateToken(c, user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	//登录成功
	loginResult := LoginResult{
		Token:        token,
		UserId:       user.UserID,
		Username:     user.Username,
		Email:        user.Email,
		UserType:     user.UserType,
		Info:         user.BasicInfo,
		Interdiction: user.Interdiction,
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "登录成功",
		"data":    loginResult,
	})

}

//修改密码
/*
	request:
	{
		"email":string,
		"password":string,
		"new_password":string,
		"new_password2":string,
	}
*/
func ResetPassword(c *gin.Context) {
	pass, user := AccountCheck(c)

	if !pass {
		return
	}

	newPassword := c.Request.FormValue("new_password")
	newPassword2 := c.Request.FormValue("new_password2")
	if newPassword != newPassword2 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "密码不一致",
		})
		return
	}

	err := server.UpdateAUserPassword(&user, newPassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "密码修改成功",
	})
	return
}

//返回用户信息
/*
	request:
	{
		"id":int,
	}
*/
func ReturnAccountInfo(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 0, 64)

	if user, notfound := server.QueryAUerById(userID); notfound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户id不存在",
			"userid":  userID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    user,
		})
	}
}

//用户个人信息修改（需要登录状态验证）
/*
	request:
	{
		"user_id":int,
		"username":string,
		"email":string,
		"info":string,
	}
*/

func ResetAccountInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*auth.CustomClaims)

	userId := claims.UserID
	username := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	info := c.Request.FormValue("info")

	var user dao.User
	var notfound bool

	if user, notfound = server.QueryAUerById(userId); notfound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户id不存在",
		})
		return
	}

	if blank, message := InputBlankCheck(username, email); blank {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": message,
		})
	}
	if duplicate := DuplicateCheck(c, userId, username, email, "ResetAccountInfo"); duplicate {
		return
	}

	if err := server.UpdateAUser(&user, username, email, info); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "用户信息已更新",
	})
	return
}

//创建收藏夹（需要进行登录状态验证）
/*
	request:
	{
		"user_id":int,
		"folder_name":string,
		"folder_info":string,
	}

*/
func CreateAFolder(c *gin.Context) {
	claims := c.MustGet("claims").(*auth.CustomClaims)
	userId := claims.UserID
	folderName := c.Request.FormValue("folder_name")
	folderInfo := c.Request.FormValue("folder_info")

	var user dao.User
	var notfound bool

	if user, notfound = server.QueryAUerById(userId); notfound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户id不存在",
		})
		return
	}

	if err, folderId := server.CreateAFolder(folderName, folderInfo, userId); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	} else {
		data := dao.Folder{
			FolderID:   folderId,
			FolderName: folderName,
			FolderInfo: folderInfo,
			UserID:     user.UserID,
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "收藏夹创建成功",
			"data":    data,
			"claims":  claims,
		})
	}
}

//添加收藏（需要登录状态验证）
/*
	request:
	{
		"user_id":int,
		"folder_id":int,
		"doc_id":int,
		"doc_info":string,
	}
*/

func AddToMyFolder(c *gin.Context) {
	claims := c.MustGet("claims").(*auth.CustomClaims)
	userId := claims.UserID
	folderId := c.Request.FormValue("folder_id")
	docId, _ := strconv.ParseUint(c.Request.FormValue("doc_id"), 0, 64)
	docInfo := c.Request.FormValue("doc_info")

	folder, notFound := server.QueryAFolderByID(folderId)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "收藏夹不存在",
		})
		return
	}

	if folder.UserID != userId {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "收藏夹拥有者与请求者不符",
		})
		return
	}

	err, _ := server.CreateAFavorite(folder.FolderID, docId, docInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "已收藏",
	})
	return
}
