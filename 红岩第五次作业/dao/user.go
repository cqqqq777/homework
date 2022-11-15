package dao

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Userdata = make(map[string]string)

// Read 将user.data的数据读取出来存放在map中
func Read() {
	var data []byte
	file, err1 := os.Open("user.data")
	if err1 != nil {
		return
	}
	defer file.Close()
	for {
		buffer := make([]byte, 2048)
		n, err2 := file.Read(buffer)
		if err2 != nil {
			if err2 == io.EOF {
				break
			} else {
				fmt.Println(err2)
			}
		}
		data = append(data, buffer[:n]...)
	}
	err3 := json.Unmarshal(data, &Userdata)
	if err3 != nil {
		return
	}
}

// Add 添加用户数据
func Add(username, password string) {
	Userdata[username] = password
}

// Check 检查map里面有没有这个用户名
func Check(username string) bool {
	_, ok := Userdata[username]
	return ok
}

// Storage 将注册后的数据存放在user.data中
func Storage() {
	sto, err1 := json.Marshal(Userdata)
	if err1 != nil {
		return
	}
	file, err2 := os.OpenFile("user.data", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer file.Close()
	if err2 != nil {

		return
	}
	_, err3 := file.Write(sto)
	if err3 != nil {
		return
	}
}
