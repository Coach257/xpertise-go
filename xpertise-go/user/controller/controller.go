package controller

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
	"xpertise-go/dao"
	auth "xpertise-go/user/auth"
	"xpertise-go/user/server"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Xpertise Scholar",
	})
}

func CreateAUser(c *gin.Context) {
	user := dao.User{Username: "Rolin", Password: "123456", Email: "1207640183@qq.com"}
	if err := server.CreateAUser(&user); err != nil {
		c.JSON(0, gin.H{"message": err})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}

func DeleteAStudentByID(c *gin.Context) {
	sid, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	server.DeleteAStudentByID(sid)
	c.JSON(200, gin.H{"message": "success"})
}

func QueryAllStudents(c *gin.Context) {
	students := server.QueryAllStudents()
	c.IndentedJSON(200, students)
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

func generateToken(c *gin.Context, user dao.User) (string, error) {
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

func Register(c *gin.Context) {
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

	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	password2 := c.Request.FormValue("password2")
	email := c.Request.FormValue("email")
	info := c.Request.FormValue("info")

	if _, notfound := server.QueryAUserByUsername(username); notfound != true {
		c.JSON(200, gin.H{"success": false, "message": "用户名已被占用"})
		return
	}
	if password != password2 {
		c.JSON(200, gin.H{"success": false, "message": "两次密码不一致"})
		return
	}
	if email == "" {
		c.JSON(200, gin.H{"success": false, "message": "未输入邮箱"})
		return
	}
	if _, notfound := server.QueryAUserByEmail(email); notfound != true {
		c.JSON(200, gin.H{"success": false, "message": "邮箱已被占用"})
		return
	}

	user := dao.User{Username: username, Password: password, Email: email, BasicInfo: info}
	server.CreateAUser(&user)
	c.JSON(200, gin.H{"success": true, "message": "用户创建成功"})
}

type LoginResult struct {
	Token        string `json:"token"`
	Userid       uint64 `json:"userid"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Usertype     int    `json:"usertype"`
	Info         string `json:"info"`
	Interdiction bool   `json:"interdiction"`
}

//验证账户信息
func AccountCheck(c *gin.Context) (bool,dao.User) {
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
		return false,user
	}

	if notfound {
		c.JSON(200, gin.H{
			"success": false,
			"message": "用户或邮箱不存在",
		})
		return false,user
	}

	if user.Password != password {
		c.JSON(200, gin.H{
			"success": false,
			"message": "用户名或密码错误",
		})
		return false,user
	}

	return true,user
}

func Login(c *gin.Context) {
	/*
		request:
		{
			"username":string,
			"email":string,
			"password":string
		}
	*/

	pass,user:=AccountCheck(c)

	if !pass {
		return
	}

	token, err := generateToken(c, user)

	// debug
	log.Println(token)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	//登录成功
	loginresult := LoginResult{
		Token:        token,
		Userid:       user.UserID,
		Username:     user.Username,
		Email:        user.Email,
		Usertype:     user.Usertype,
		Info:         user.BasicInfo,
		Interdiction: user.Interdiction,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "登录成功",
		"data":    loginresult,
	})

}

func ResetPassword(c *gin.Context) {
	/*
		request:
		{
			"email":string,
			"password":string,
			"newpassword":string,
			"newpassword2":string,
		}
	*/
	pass,user :=AccountCheck(c)

	if !pass{
		return
	}

	newpassword := c.Request.FormValue("newpassword")
	newpassword2 := c.Request.FormValue("newpassword2")

	if newpassword!=newpassword2{
		c.JSON(http.StatusOK,gin.H{
			"success":false,
			"message":"密码不一致",
		})
		return
	}

	err:=server.UpdateAUserPassword(&user,newpassword)

	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"success":false,
			"message":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"message":"密码修改成功",
	})

	return
}
