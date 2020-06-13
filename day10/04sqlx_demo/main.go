package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" // init()
)

// go连接MySQL实例

var db *sqlx.DB // 是一个连接池对象

func initDB() (err error) {
	// 连接数据库
	// 用户名：密码@tcp(ip：端口)/数据库名称
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	// open不会校验用户名和密码是否正确
	db, err = sqlx.Connect("mysql", dsn) // 此处不用 := ,调用函数外部声明的全局变量db
	if err != nil {                      // dsn格式不正确的时候报错
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大闲置连接数
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	Name string
	Age  int
	ID   int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed:%v\n", err)
		return
	}
	sqlStr1 := `select * from user;`
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := `select * from user;`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
		return
	}
	fmt.Printf("usrList:%#v\n", userList)
}
