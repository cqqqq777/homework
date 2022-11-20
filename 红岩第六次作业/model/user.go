package model

type User struct {
	ID       int
	Username string
	Password string
	Secure   string
}
type Message struct {
	MesObj  string //留言对象
	Content string //留言内容
	MesPer  string //留言人、
	ID      string //留言编号
}
