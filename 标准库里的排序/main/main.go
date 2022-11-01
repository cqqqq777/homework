// 随机生成一百个数放进切片然后排序
package main

import (
	"fmt"
	"math/rand"
	sort2 "sort"
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
	sort2.Ints(sort)
	for _, val := range sort {
		fmt.Printf("%d ", val)
	}
}
