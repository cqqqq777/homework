package main

import (
	"fmt"
)

var x int

func main() {
	//一个管道用来让x一直累加，另一个管道用来阻塞主线程
	exitCh := make(chan bool)
	ch := make(chan int, 100)
	go func(ch chan int, e chan bool) {
		for i := 0; i < 100000; i++ {
			x += <-ch
		}
		exitCh <- true
	}(ch, exitCh)
	go func(ch chan int) {
		for i := 0; i < 100000; i++ {
			ch <- 1
		}
	}(ch)
	<-exitCh
	fmt.Println(x)
}
