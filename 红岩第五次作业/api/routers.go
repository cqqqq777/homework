package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.POST("/enroll", enroll)
	r.POST("/login", login)
	r.Run()
}
