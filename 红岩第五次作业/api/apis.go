package api

import (
	"github.com/gin-gonic/gin"
	"go_code/gin/dao"
	"net/http"
)

// 注册
func enroll(c *gin.Context) {
	// 传入用户名和密码
	dao.Read()
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证用户名是否重复
	ok := dao.Check(username)
	// 重复则退出
	if ok {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Repeat registration",
		})
		return
	}
	//将新用户信息保存下来
	dao.Add(username, password)
	dao.Storage()
	// 以 JSON 格式返回信息
	c.JSON(http.StatusOK, gin.H{
		"message": "add user successful",
	})
}

// 登录
func login(c *gin.Context) {
	dao.Read()
	username := c.PostForm("username")
	password := c.PostForm("password")
	ok := dao.Check(username)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Not registered",
		})
		return
	}
	if dao.Userdata[username] != password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Wrong password",
		})
		return
	}
	//设置cookie
	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successful login"})
}
