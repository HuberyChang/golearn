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
	db, err = sql.Open("mysql", dsn) // 此处不用 := ,调用全局变量db
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

func queryOne(id int) {
	var u1 user
	// 1、单条查询
	sqlStr := `select id, name, age from user where id=?;`
	// for i := 0; i < 11; i++ {
	// 	fmt.Println(i)
	// 	db.QueryRow(sqlStr, 1) // 从连接池里拿一个连接出来去数据库查询单挑记录
	// }
	rowObj := db.QueryRow(sqlStr, id)
	rowObj.Scan(&u1.id, &u1.name, &u1.age) // 必须对rowObj对象调用scan函数，因为scan函数自带释放连接函数
	fmt.Printf("u1:%#v\n", u1)
}

func insert() {

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed: %v\n", err)
	}
	fmt.Println("数据库连接成功！")
	queryOne(1)
}
