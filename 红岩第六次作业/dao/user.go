package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_code/gin/model"
)

var db *sql.DB

func InitDB() error {
	dsn := "root:zhaoxijun7@tcp(127.0.0.1:3306)/user"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// Insert 添加用户数据
func Insert(username, password, secure string) {
	db.Exec("insert into userinfo(username,password,secure) values(?,?,?)", username, password, secure)
}

// Query 检查数据库里面有没有这个用户名
func Query(username string) bool {
	rows, err := db.Query("select username from userinfo")
	if err != nil {
		return true
	}
	defer rows.Close()
	var u model.User
	//找数据库有没有这个用户，有的话返回true，没有返回false
	for rows.Next() {
		rows.Scan(&u.Username)
		if username == u.Username {
			return true
		}
	}
	return false
}

// VerifyPassword 验证用户密码是否正确
func VerifyPassword(username, password string) bool {
	row := db.QueryRow("select password from userinfo where username =?", username)
	var u model.User
	err := row.Scan(&u.Password)
	if err != nil {
		return false
	}
	return password == u.Password
}

// VerifySecure 验证密保是否正确
func VerifySecure(username, secure string) bool {
	row := db.QueryRow("select secure from userinfo where username=?", username)
	var u model.User
	err := row.Scan(&u.Secure)
	if err != nil {
		return false
	}
	return secure == u.Secure
}

// Reset 重置密码,并返回是否重置成功
func Reset(username, newPassword string) bool {
	_, err := db.Exec("update userinfo set password =?where username=?", newPassword, username)
	if err != nil {
		return false
	}
	return true
}

// Message 留言
func Message(username, content, MesObj string) bool {
	_, err := db.Exec("insert into message(content,mesobj,mesper) values(?,?,?)", content, MesObj, username)
	if err != nil {
		return false
	}
	return true
}

// Inquire 查询留言内容与留言人
func Inquire(username string) (string, string) {
	//获取用户名对应的留言内容与留言人
	row1 := db.QueryRow("select content from mesaage where mesobj=?", username)
	row2 := db.QueryRow("select content from mesaage where mesobj=?", username)
	var m model.Message
	//把获取的数据写进结构体并返回
	err := row1.Scan(&m.Content)
	if err != nil {
		return "failed", "failed"
	}
	err = row2.Scan(&m.MesPer)
	return m.Content, m.MesPer
}
