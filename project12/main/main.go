package main

import (
	"fmt"
	"strconv"
	"strings"
)

type aaaa struct {
	array []string
}

func (s *aaaa) overturning() {
	var a []int64
	var length, num int
	var str string
	fmt.Println("请输入数组长度与你想翻转的元素个数：")
	_, err := fmt.Scanln(&length, &num)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("请输入%d个整数(逗号隔开)：\n", length)
	_, err2 := fmt.Scanln(&str)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	s.array = strings.Split(str, ",")
	for _, val := range s.array {
		val1, _ := strconv.ParseInt(val, 10, 0)
		a = append(a, val1)
	}
	for i := 0; i < num/2; i++ {
		a[i], a[num-1-i] = a[num-1-i], a[i]
	}
	for _, val2 := range a {
		fmt.Printf("%d ", val2)
	}
}
func main() {
	var s aaaa
	s.overturning()
}
