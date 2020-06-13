package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

// go连接MySQL实例

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	// 连接数据库
	// 用户名：密码@tcp(ip：端口)/数据库名称
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	// open不会校验用户名和密码是否正确
	db, err = sql.Open("mysql", dsn) // 此处不用 := ,调用函数外部声明的全局变量db
	if err != nil {                  // dsn格式不正确的时候报错
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大闲置连接数
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	name string
	age  int
	id   int
}

func transactionDemo() {
	// 1. 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed:,err:%v\n", err)
	}
	// 执行多条语句
	sqlStr1 := `update user set age=age-2 where id=1`
	sqlStr2 := `update user set age=age+2 where id=2`
	// 依次执行两条sql
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql1出错，需要回滚")
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql2出错，需要回滚")
		return
	}
	// 上面两部都执行成功，提交事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交出错，需要回滚")
		return
	}
	fmt.Println("事务成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed: %v\n", err)
	}
	fmt.Println("数据库连接成功！")
	transactionDemo()
}
