package main

import (
	"fmt"
	"runtime"
	"time"
)

// 设置全局变量来堵塞主线程或者删除某个定时器
var (
	Li  = make(chan int, 1)
	Get = make(chan int, 1)
	Wu  = make(chan int, 1)
)

func main() {
	var choose int
	go WuHu()
	go Library()
	go GetUp()
	for {
		fmt.Println("请选择你想要进行的操作")
		fmt.Println("1.增加定时器")
		fmt.Println("2.删除内置定时器")
		fmt.Println("3.退出程序")
		fmt.Scan(&choose)
		switch choose {
		case 1:
			add()
		case 2:
			delete1()
		case 3:
			return
		}
	}
}

// Library 每天六点输出一句话
func Library() {
	defer fmt.Println("您已成功删除该定时器")
	var ticker2 *time.Ticker
	for {
		if len(Li) == 1 {
			runtime.Goexit()
		}
		if time.Now().Hour() == 6 {
			ticker2 = time.NewTicker(time.Hour * 24)
			break
		}
	}
	for {
		select {
		case <-ticker2.C:
			fmt.Println("我要去图书馆开卷！")
		case <-Li:
			runtime.Goexit()
		}
	}
}

// GetUp 每天四点输出一句话
func GetUp() {
	defer fmt.Println("您已成功删除该定时器")
	var ticker1 *time.Ticker
	for {
		if len(Get) == 1 {
			runtime.Goexit()
		}
		if time.Now().Hour() == 4 {
			ticker1 = time.NewTicker(time.Hour * 24)
			break
		}
	}
	for {
		select {
		case <-Get:
			runtime.Goexit()
		case <-ticker1.C:
			fmt.Println("我还能再战4小时！")
		}
	}
}

// WuHu 每隔半分钟输出“起飞”
func WuHu() {
	defer fmt.Println("您已成功删除该定时器")
	ticker := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-Wu:
			runtime.Goexit()
		case <-ticker.C:
			fmt.Println("芜湖！起飞！")
		}
	}
}

// 新增定时器
func add() {
	var b, hour, min, sec int
	var event string
	fmt.Println("请选择你想要增加的定时器类型")
	fmt.Println("1.一次性定时器")
	fmt.Println("2.重复性定时器")
	fmt.Scan(&b)
	fmt.Println("请输入你想要的定时器时长（格式：1 2 3代表一小时两分钟三秒）：")
	fmt.Scan(&hour, &min, &sec)
	switch b {
	case 1:
		fmt.Println("请设置事件(如：”起床“)")
		fmt.Scan(&event)
		go func() {
			timer := time.NewTimer(time.Second * time.Duration(sec+min*60+hour*60*60))
			<-timer.C
			fmt.Println(event)
		}()
	case 2:
		fmt.Println("请设置事件(如：”起床“)")
		fmt.Scan(&event)
		go func() {
			ticker := time.NewTicker(time.Second * time.Duration(sec+min*60+hour*60*60))
			for {
				<-ticker.C
				fmt.Println(event)
			}
		}()
	}
}

// 删除内置定时器
func delete1() {
	var b int
	fmt.Println("请选择你想要删除的定时器")
	fmt.Println("1.图书馆定时器")
	fmt.Println("2.卷王定时器")
	fmt.Println("3.芜湖定时器")
	fmt.Scan(&b)
	switch b {
	case 1:
		Li <- 1
	case 2:
		Get <- 1
	case 3:
		Wu <- 1
	}
}
