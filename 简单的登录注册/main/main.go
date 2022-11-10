package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var userdata = make(map[string]string)

func main() {
	var a int
	read()
lab1:
	fmt.Println("请选择你想要进行的操作；")
	fmt.Println("1.登录已有账户")
	fmt.Println("2.注册新账户")
	fmt.Println("3.退出程序")
	fmt.Scan(&a)
	switch a {
	case 1:
		login()
		goto lab1
	case 2:
		enroll()
		goto lab1
	case 3:
		storage()
		return
	}
}

// 登录
func login() {
	var userName, Pwd string
	var a int
lab1:
	fmt.Println("请输入用户名；")
	fmt.Scan(&userName)
	_, ok := userdata[userName]
	if !ok {
		fmt.Println("系统中不存在此用户,请重新输入用户名")
		goto lab1
	}
lab2:
	fmt.Println("请输入密码")
	fmt.Scan(&Pwd)
	if userdata[userName] == Pwd {
		fmt.Println("登录成功！！！")
	lab3:
		fmt.Println("请选择接下来的操作")
		fmt.Println("1.修改密码")
		fmt.Println("2.返回上一级")
		fmt.Scan(&a)
		switch a {
		case 1:
			revise(userName, Pwd)
			if revise(userName, Pwd) {
				goto lab3
			}
		case 2:
			return
		}
	} else {
		fmt.Println("密码错误,请重新输入")
		goto lab2
	}
}

// 打开程序时先将用户数据读入map中
func read() {
	var data []byte
	file, err1 := os.Open("user.data")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer file.Close()
	for {
		buffer := make([]byte, 2048)
		n, err2 := file.Read(buffer)
		if err2 != nil {
			if err2 == io.EOF {
				break
			} else {
				fmt.Println(err2)
			}
		}
		data = buffer[:n]
	}
	err3 := json.Unmarshal(data, &userdata)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
}

// 注册
func enroll() {
	var userName, Pwd string
lab1:
	fmt.Println("请输入新用户名:")
	fmt.Scan(&userName)
	_, ok := userdata[userName]
	if ok {
		fmt.Println("已存在此用户名，请重新输入")
		goto lab1
	}
lab2:
	fmt.Println("请输入密码(至少8位):")
	fmt.Scan(&Pwd)
	if len(Pwd) < 8 {
		fmt.Println("密码至少8位!!!")
		goto lab2
	}
	userdata[userName] = Pwd
	fmt.Println("注册成功")
}

// 修改密码
func revise(u, p string) bool {
	var a, b, c string
	fmt.Println("请输入原密码")
	fmt.Scan(&a)
	if a != p {
		fmt.Println("密码错误!!!")
		return true
	}
	fmt.Println("请输入新密码（至少8位）：")
	fmt.Scan(b)
	if len(b) < 8 {
		fmt.Println("密码至少8位！！！")
		return true
	}
	fmt.Println("请再次输入新密码：")
	fmt.Scan(&c)
	if c != b {
		fmt.Println("两次密码输入不一致")
		return true
	}
	userdata[u] = b
	fmt.Println("修改密码成功!")
	return false
}
func storage() {
	sto, err1 := json.Marshal(userdata)
	if err1 != nil {
		fmt.Println("序列化失败：", err1)
	}
	file, err2 := os.OpenFile("user.data", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err2 != nil {
		fmt.Println("保存数据失败：", err2)
	}
	defer file.Close()
	_, err3 := file.Write(sto)
	if err3 != nil {
		fmt.Println("保存数据失败失败：", err3)
	}
}
