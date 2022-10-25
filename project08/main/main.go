// 简易计算器
package main
//计算器函数在另一个包里，这里引入另一个包
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
//定义一个空切片以接受每次计算结果，然后输出切片
	r := make([]float64, 0, 10000)
	for {
		fmt.Println("请输入一个算式（用空格隔开）：")
		fmt.Scanln(&a, &cmd, &b)
		result := main01.Cla(a, b, cmd)
		//fmt.Println(result)
		r = append(r, result)
		fmt.Println(r)
	}
}
