package main

import (
	"fmt"
	"sort"
)

type Interface interface {
	//Len 方法返回集合中的元素个数
	Len() int
	//Less 方法报告索引i的元素是否比索引j的元素小
	Less(i, j int) bool
	//Swap 方法交换索引i和j的两个元素
	Swap(i, j int)
}

type Student struct {
	Name   string
	StuNum int
}

// 自定义“学生”切片类型存储学生数据
type database []Student

//让database实现接口,以调用sort包中的函数进行排序

func (d database) Len() int {
	return len(d)
}

func (d database) Less(i, j int) bool {
	if d[i].StuNum < d[j].StuNum {
		return true
	} else {
		return false
	}
}

func (d database) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func sort1() {
	data := database{
		{"小明", 2022192},
		{"小红", 2022999},
		{"小刚", 2022544},
		{"小美", 2022098},
	}
	//输出原始顺序的姓名和学号
	for i := 0; i < len(data); i++ {
		fmt.Println("姓名:"+data[i].Name, "学号:", data[i].StuNum)
	}
	fmt.Println()
	//使用标准库中的sort.Stable函数按学号排序
	sort.Stable(data)
	//输出按学号升序排序后的姓名和学号
	for i := 0; i < len(data); i++ {
		fmt.Println("姓名:"+data[i].Name, "学号:", data[i].StuNum)
	}
}
