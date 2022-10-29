package main

import "fmt"

var (
	name string
	pwd  string
)

func main() {
	fmt.Println("请输入用户名：")
	fmt.Scanln(&name)
	fmt.Println("请输入密码：")
	fmt.Scanln(&pwd)
	if name == "abc" && pwd == "abc123" {
		fmt.Println("登录成功")
	} else {
		fmt.Println("密码错误")
	}
}
