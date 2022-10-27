package main

import (
	"fmt"
)

type Interface interface {
	call()
}
type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (d Dog) call() {
	fmt.Println(d.name + "会汪汪叫。")
}
func (c Cat) call() {
	fmt.Println(c.name + "会喵喵叫。")
}
func animals(i Interface) {
	i.call()
}
func homework() {
	dog := Dog{
		name: "tom",
	}
	cat := Cat{
		name: "jack",
	}
	animals(dog)
	animals(cat)
}
