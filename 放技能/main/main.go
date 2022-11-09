package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 抽象出skill的属性
type skill struct {
	Name  string
	Words string
}

// 用一个map当作技能库
var skillMap = make(map[int]skill)

// Add 添加技能到技能库
func Add() {
	var Skill skill
	var name, words string
	var num int
	fmt.Println("请输入你想添加的技能名称并为其编号（先输入技能名称后输入编号，之间使用空格或换行符隔开，技能编号不小于2）：")
	_, err := fmt.Scan(&name, &num)
	if err != nil {
		fmt.Println("输入失败")
		return
	}
	if senWords(name) {
		fmt.Println("警告：请不要包含敏感词")
		return
	}
	fmt.Println("请输入你想要与此技能绑定的模板（如“尝尝我的厉害吧!!!”）：")
	_, err2 := fmt.Scan(&words)
	if err2 != nil {
		fmt.Println("输入失败")
		return
	}
	if senWords(words) {
		fmt.Println("警告：请不要包含敏感词")
		return
	}
	Skill.Name = name
	Skill.Words = words
	skillMap[num] = Skill
}
func main() {
	for {
		var choose int
		fmt.Println("请选择你想要进行的操作：")
		fmt.Println("1.释放技能库的技能")
		fmt.Println("2.添加技能到技能库")
		fmt.Println("3.退出")
		_, err := fmt.Scan(&choose)
		if err != nil {
			fmt.Println("输入失败")
			continue
		}
		switch choose {
		case 1:
			system()
		case 2:
			Add()
		case 3:
			return
		default:
			fmt.Println("请按要求输入数字")
		}
	}
}

// ReleaseSkill 放技能
func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}

// 初始化技能库
func init() {
	skillMap[1] = skill{
		Name:  "龙卷风摧毁停车场",
		Words: "尝尝我的厉害吧!!!",
	}
	skillMap[2] = skill{
		Name:  "乌鸦坐飞机",
		Words: "没人可以在我的手中活下来!!!",
	}
}

// 放技能库的技能
func system() {
	var a int
	fmt.Println("请选择你想要释放的技能（输入其编号）：")
	for k, v := range skillMap {
		fmt.Println(strconv.FormatInt(int64(k), 10) + ".技能名称：" + v.Name + " 模板:" + v.Words)
	}
	_, err := fmt.Scan(&a)
	if err != nil {
		return
	}
	ReleaseSkill(skillMap[a].Name, func(s string) {
		fmt.Println(skillMap[a].Words, skillMap[a].Name)
	})
}

// 判断敏感词
func senWords(s string) bool {
	var b bool
	senWord := make([]string, 0)
	senWord = append(senWord, "傻逼")
	for i := 0; i < len(senWord); i++ {
		b = strings.Contains(s, senWord[i])
	}
	return b
}
