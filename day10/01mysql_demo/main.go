package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

// go连接MySQL实例

var db *sql.DB // 是一个连接池对象

func query() {

}

func insert() {

}

func main() {
	// 连接数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/goday10"
	// 连接数据库
	db, err := sql.Open("mysql", dsn) // open不会校验用户名和密码是否正确
	if err != nil {                   // dsn格式不正确的时候报错
		fmt.Printf("open %s invalid, err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}
	fmt.Println("success")
}
