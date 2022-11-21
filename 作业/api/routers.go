package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.POST("/enroll", enroll)
	r.POST("/login", login)
	r.POST("/revise", revise)
	r.GET("/sayHello", check, sayHello)
	//留言功能先检查是否登录，检查cookie
	group := r.Group("/MessageBoards", check)
	{
		group.GET("/inquire", inquire)
		group.POST("/message", message)
	}
	err := r.Run()
	if err != nil {
		return
	}
}
