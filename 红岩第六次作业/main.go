package main

import (
	"fmt"
	"go_code/gin/api"
	"go_code/gin/dao"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		fmt.Println("link failed")
		return
	}
	api.InitRouter()
}
