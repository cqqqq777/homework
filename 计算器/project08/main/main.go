// 简易计算器
package main

import (
	"fmt"
	"go_code/project08/main01"
)

func main() {
	var (
		a   float64
		b   float64
		cmd string
	)
	r := make([]float64, 0, 10000)
	for {
		fmt.Println("请输入一个算式（用空格隔开）：")
		_, _ = fmt.Scanln(&a, &cmd, &b)
		result := main01.Cla(a, b, cmd)
		r = append(r, result)
		fmt.Println(r)
	}
}
