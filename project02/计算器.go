package main

import "fmt"

var (
	a int
	b string
	c int
)

func main() {
	fmt.Scanln(&a, &b, &c)
	switch b {
	case "+":
		fmt.Println(a, b, c, "=", a+c)
	case "-":
		fmt.Println(a, b, c, "=", a-c)
	case "*":
		fmt.Println(a, b, c, "=", a*c)
	case "/":
		fmt.Println(a, b, c, "=", float32(a)/float32(c))
	case "%":
		fmt.Println(a, b, c, "=", a%c)
	default:
		fmt.Println("error")
	}
}
