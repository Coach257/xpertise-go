package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "buaa21xpertise"
)

type CustomClaims struct {
	UserID 	 uint64 `json:"user_id"`
	Username string `json:"username"`
	Email	 string `json:"email"`
	jwt.StandardClaims
}

//获取signkey
func GetSignKey() string{
	return SignKey
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 生成token
func (j *JWT)CreateToken(claims CustomClaims)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(j.SigningKey)
}

// 解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims,error){
	token, err :=jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey,nil
	})
	if err !=nil{
		if ve,ok := err.(*jwt.ValidationError);ok{
			if ve.Errors&jwt.ValidationErrorMalformed != 0{
				return nil, TokenMalformed
			}else if ve.Errors & jwt.ValidationErrorExpired !=0{
				return nil, TokenExpired
			}else if ve.Errors & jwt.ValidationErrorNotValidYet != 0{
				return nil, TokenNotValidYet
			}else{
				return nil, TokenInvalid
			}
		}
	}
	if claims,ok := token.Claims.(*CustomClaims);ok && token.Valid{
		return claims,nil
	}
	return nil, TokenInvalid
}


// jwt中间件，用于检查token
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试获取请求头部的token
		token := c.Request.Header.Get("token")
		//token:=c.Request.FormValue("token")

		// 如果token为空，那就直接return
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		// 如果获取到了token

		// 新建一个jwt实例
		j := NewJWT()

		// 解析token包含的信息，返回给claims
		claims,err := j.ParseToken(token)

		// 如果解析报错
		if err !=nil{
			// token过期了
			if err == TokenExpired {
				c.JSON(http.StatusOK,gin.H{
					"success":false,
					"message":"token已过期",
				})
				c.Abort()
				return
			}

			// 其他错误
			c.JSON(http.StatusOK,gin.H{
				"success":false,
				"message":err.Error(),
			})
			c.Abort()
			return
		}

		// 如果没有报错，比对请求者id与token中id是否一致
		userIdFromToken := claims.UserID
		userId,_:=strconv.ParseUint(c.Request.FormValue("user_id"),0,64)
		if userId!=userIdFromToken{
			c.JSON(http.StatusOK,gin.H{
				"success":false,
				"message":"请求用户与token记录的用户不一致",
			})
			c.Abort()
			return
		}

		//没有任何问题，就将claims中的信息传给下一个handler
		c.Set("claims",claims)
		c.Next()
	}
}

