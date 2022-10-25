package main

import "fmt"

// Cat 定义“猫”结构体，包括猫的名字，体重，品种，颜色与性别
type Cat struct {
	Name                 string
	Weight               float64
	Gender, Color, Breed string
}

// 创建一个Cat的方法描述猫的行为
func (a Cat) attribute() {
	fmt.Println(a.Name, "会喵喵叫，喜欢玩球，喜欢睡觉，爱干净。")
}

func main() {
	//为tom添加属性
	tom := Cat{
		Name:   "tom",
		Weight: 5.6,
		Gender: "雄",
		Color:  "黄色",
		Breed:  "短毛猫",
	}
	fmt.Println("我的猫叫tom，它的体重:", tom.Weight, "kg 性别:", tom.Gender, " 颜色:", tom.Color, "品种：", tom.Breed)
	tom.attribute()
}
