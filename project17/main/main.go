package main

import (
	"fmt"
	"os"
)

func main() {
	//创建文件
	a, _ := os.Create("plan.txt")
	//向文件里面写
	_, _ = a.WriteString("I’m not afraid of difficulties and insist on learning programming")
	b := make([]byte, 100, 100)
	//打开文件
	file, _ := os.Open("plan.txt")
	//读文件，用n接受返回的字节数
	n, _ := file.Read(b)
	//将读到的字节转化成string类型并打印
	fmt.Println(string(b[:n]))
	//关闭文件
	_ = file.Close()
}
