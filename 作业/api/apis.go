package api

import (
	"github.com/gin-gonic/gin"
	"go_code/gin/dao"
	"net/http"
)

// 注册
func enroll(c *gin.Context) {
	// 传入用户名和密码和密保
	username := c.PostForm("username")
	password := c.PostForm("password")
	secure := c.PostForm("secure")
	// 验证用户名是否重复
	ok := dao.Query(username)
	// 重复则退出
	if ok {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Repeat registration",
		})
		return
	}
	//将新用户信息保存下来
	dao.Insert(username, password, secure)
	// 以 JSON 格式返回信息
	c.JSON(http.StatusOK, gin.H{
		"message": "add user successful",
	})
}

// 登录
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证是否注册
	ok := dao.Query(username)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Not registered",
		})
		return
	}
	// 验证密码
	if !dao.VerifyPassword(username, password) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Wrong password",
		})
		return
	}
	//设置cookie
	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successful login"})
}

// 检查用户是否登录
func check(c *gin.Context) {
	_, err := c.Cookie("gin_cookie")
	if err != nil {
		c.JSON(http.StatusOK, "Not logged in")
		c.Abort()
		return
	}
	c.Next()
}

// 通过密保修改密码
func revise(c *gin.Context) {
	username := c.PostForm("username")
	newPassword := c.PostForm("newPassword")
	secure := c.PostForm("secure")
	//验证密保
	if dao.VerifySecure(username, secure) {
		if dao.Reset(username, newPassword) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "wrong answer"})
	}
}

// 查看留言
func inquire(c *gin.Context) {
	username := c.Query("username")
	content, MesPer := dao.Inquire(username)
	c.JSON(http.StatusOK, gin.H{
		"content":       content,
		"MessagePerson": MesPer,
	})
}

// 留言
func message(c *gin.Context) {
	username := c.Query("username")
	content := c.PostForm("content")
	MesObj := c.PostForm("MesObj")
	ok := dao.Message(username, content, MesObj)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
	}
}

// 说你好
func sayHello(c *gin.Context) {
	_, err := c.Cookie("gin_cookie")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "游客你好"})
	} else {
		username := c.Query("username")
		c.JSON(http.StatusOK, gin.H{
			"message": username + "你好",
		})
	}
}
