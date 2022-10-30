package main

import "fmt"

func main() {
	//创建两个管道一个接受奇数一个接受偶数
	odd := make(chan int, 50)
	even := make(chan int, 50)
	//发送奇数
	go func(odd chan int) {
		for i := 1; i <= 50; i++ {
			odd <- 2*i - 1
		}
	}(odd)
	//发送偶数
	go func(even chan int) {
		for i := 1; i <= 50; i++ {
			even <- 2 * i
		}
	}(even)
	//轮流打印
	for i := 1; i <= 50; i++ {
		fmt.Printf("%d,%d,", <-odd, <-even)
	}
}
