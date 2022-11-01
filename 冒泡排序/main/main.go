// 冒泡排序
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sort := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		sort = append(sort, rand.Intn(1000))
	}
	for _, val := range sort {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	for i := 1; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			if sort[j] > sort[1+j] {
				sort[j], sort[j+1] = sort[j+1], sort[j]
			}
		}
	}
	for _, val := range sort {
		fmt.Printf("%d ", val)
	}
}
