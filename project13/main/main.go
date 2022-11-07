package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type aa []int64

func (d aa) Len() int { return len(d) }

func (d aa) Less(i, j int) bool { return d[i] < d[j] }

func (d aa) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

func main() {
	var str string
	var a aa
	fmt.Println("请输入想排序的一组整数(用逗号隔开)：")
	_, err := fmt.Scanln(&str)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		str1 := strings.Split(str, ",")
		for _, val := range str1 {
			b, err1 := strconv.ParseInt(val, 10, 0)
			if err1 != nil {
				fmt.Println(err1)
				return
			} else {
				a = append(a, b)
			}
		}
	}
	sort.Stable(a)
	for _, val1 := range a {
		fmt.Printf("%d ", val1)
	}
}
