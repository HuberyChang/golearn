package main

// sql注入

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

func sqlInjectDemo(name string)  {
	// 自己拼接sql语句的字符串
	sqlStr:=fmt.Sprintf("select * from user where name='%s'",name)
	fmt.Printf("sql:%s\n",sqlStr)

	var users []user
	err:=db.Select(&users,sqlStr)
	if err != nil {
		fmt.Printf("exec failed,err:%v\n",err)
		return
	}
	for _,u:=range users{
		fmt.Printf("user:%#v\n",u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed:%v\n", err)
		return
	}
	// sql注入的几种实例
	sqlInjectDemo("交罚款了")
}