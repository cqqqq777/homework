// package main
//
// import (
//
//	"fmt"
//
// )
//
//	type Interface interface {
//		call()
//	}
//
//	type Dog struct {
//		name string
//	}
//
//	type Cat struct {
//		name string
//	}
//
//	func (d Dog) call() {
//		fmt.Println(d.name + "会汪汪叫。")
//	}
//
//	func (c Cat) call() {
//		fmt.Println(c.name + "会喵喵叫。")
//	}
//
//	func animals(i Interface) {
//		i.call()
//	}
//
//	func homework() {
//		dog := Dog{
//			name: "tom",
//		}
//		cat := Cat{
//			name: "jack",
//		}
//		animals(dog)
//		animals(cat)
//	}
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func main() {
	people := ByAge{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Sort(people)
	fmt.Println(people)
	// Output:
	// [Bob: 31 John: 42 Michael: 17 Jenny: 26]
	// [Michael: 17 Jenny: 26 Bob: 31 John: 42]
}
