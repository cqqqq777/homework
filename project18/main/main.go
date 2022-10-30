package main

import "fmt"

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		//打印的速度远慢于判断，循环开了goroutine后在打印前i的值已经变化了
		go func() {
			fmt.Println(i)
		}()
		//i等于9时因为over没有缓存，并且没有其他协程来接受over里面的值，主线程直接阻塞了不会执行下一步，导致死锁
		if i == 9 {
			over <- true
		}
	}
	<-over
	fmt.Println("over!!!")
}

//解决
func solve() {
	intCh := make(chan int, 10)
	over := make(chan bool)
	go func(intCh chan int, over chan bool) {
		for v := range intCh {
			fmt.Println(v)
		}
		over <- true
	}(intCh, over)
	for i := 0; i < 10; i++ {
		intCh <- i
	}
	close(intCh)
	<-over
	fmt.Println("over!!!")
}
