package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go_code/gin/model"
	"net/http"
	"strings"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var Secret = []byte("zxj")

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		c.JSON(http.StatusOK, gin.H{"message": "empty header"})
		c.Abort()
		return
	}
	parts := strings.SplitN(header, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"msg": "wrong format",
		})
		c.Abort()
		return
	}
	TokenString := parts[1]
	mc, err := ParseToken(TokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Invalid Token",
		})
		c.Abort()
		return
	}
	// 将当前请求的username信息保存到请求的上下文c上
	c.Set("username", mc.Username)
	c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
}

// GenToken 生成token
func GenToken(username string) string {
	//创建claim实例
	c := model.MyClaims{
		Username: username, StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "localhost",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(Secret)
	if err != nil {
		return ""
	}
	return tokenString
}

// ParseToken 解析token
func ParseToken(TokenString string) (*model.MyClaims, error) {
	var mc = new(model.MyClaims)
	token, err := jwt.ParseWithClaims(TokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
