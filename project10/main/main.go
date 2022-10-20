// 判断一句话是否是回文串
package main

import "fmt"

func main() {
	var sen string
	//定义sum最后判断是否为回文串
	sum := 0
	fmt.Println("请输入一句话：")
	fmt.Scanln(&sen)
	//制作一个空切片存放输入的句子
	r := make([]rune, 0, len(sen))
	//用range遍历将句子中的每一个字放入切片中
	for _, v := range sen {
		r = append(r, v)
	}
	//判断第n个字与倒数第n个字是否相同，若相同，sum值+1，句子有多少个字就判断多少次
	for i := 0; i < len(r); i++ {
		if r[i] == r[len(r)-1-i] {
			sum += 1
		}
	}
	//若最终sum的值等于字的个数，就是回文串
	if sum == len(r) {
		fmt.Println("这句话是回文串")
	} else {
		fmt.Println("这句话不是回文串")
	}
}
