package api

import (
	"github.com/gin-gonic/gin"
	"go_code/gin/api/middleware"
	"go_code/gin/dao"
	"go_code/gin/utils"
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
		utils.RespFail(c, "repeat register")
		return
	}
	//将新用户信息保存下来
	dao.Insert(username, password, secure)
	// 以 JSON 格式返回信息
	utils.RespSuccess(c, "add user successful")
}

// 登录
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证是否注册
	ok := dao.Query(username)
	if !ok {
		utils.RespFail(c, "Not registered")
		return
	}
	// 验证密码
	if !dao.VerifyPassword(username, password) {
		utils.RespFail(c, "wrong password")
		return
	}
	TokenString := middleware.GenToken(username)
	c.JSON(http.StatusOK, gin.H{"message": "Successful login", "token": TokenString})
}

// 通过密保修改密码
func revise(c *gin.Context) {
	username := c.PostForm("username")
	newPassword := c.PostForm("newPassword")
	secure := c.PostForm("secure")
	//验证密保
	if dao.VerifySecure(username, secure) {
		if dao.Reset(username, newPassword) {
			utils.RespSuccess(c, "success")
		}
	} else {
		utils.RespFail(c, "wrong answer")
	}
}

// 查看留言
func inquire(c *gin.Context) {
	username := c.MustGet("username").(string) //类型断言
	content, MesPer := dao.Inquire(username)
	c.JSON(http.StatusOK, gin.H{
		"content":       content,
		"MessagePerson": MesPer,
	})
}

// 留言
func message(c *gin.Context) {
	username := c.MustGet("username").(string)
	content := c.PostForm("content")
	MesObj := c.PostForm("MesObj")
	ok := dao.Message(username, content, MesObj)
	if ok {
		utils.RespSuccess(c, "success")
	} else {
		utils.RespSuccess(c, "failed")
	}
}

// 说你好
func sayHello(c *gin.Context) {
	_, ok := c.Get("username")
	if !ok {
		utils.RespSuccess(c, "游客你好")
	} else {
		username := c.Query("username")
		utils.RespSuccess(c, username+"你好")
	}
}
