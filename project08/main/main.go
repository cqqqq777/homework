// 简易计算器
package main

import (
	"fmt"
	"go_code/project08/main01"
)

func main() {
	var (
		a        float64
		b        float64
		Operator string
	)
	fmt.Println("请输入一个算式：")
	fmt.Scanln(&a, &Operator, &b)
	result := main01.Cla(a, b, Operator)
	fmt.Println(result)
}
