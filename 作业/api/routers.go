package api

import (
	"github.com/gin-gonic/gin"
	"go_code/gin/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.POST("/enroll", enroll)
	r.POST("/login", login)
	r.POST("/revise", revise)
	r.GET("/sayHello", middleware.JWTAuthMiddleware, sayHello)
	//使用留言功能前先鉴权，这里使用的是JWT
	group := r.Group("/MessageBoards", middleware.JWTAuthMiddleware)
	{
		group.GET("/inquire", inquire)
		group.POST("/message", message)
	}
	err := r.Run()
	if err != nil {
		return
	}
}
