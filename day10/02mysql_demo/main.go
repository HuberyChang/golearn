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

// 查询单条
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

// 查询多行
func queryMore(n int) {
	// 1. sql语句
	sqlStr := `select id,name,age from user where id>?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%#v\n", sqlStr, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

// 插入数据
func insert() {
	// 1. 写sql语句
	sqlStr := `insert into user(name,age) values("发丝",79)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed:%v\n", err)
		return
	}
	// 如果是插入数据的操作，能够拿到插入数据的ID
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get last id failed:%v\n", err)
		return
	}
	fmt.Printf("id:%v\n", id)
}

// 更新操作
func updateRow(newAge, id int) {
	sqlStr := `update user set age=? where id=?;`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rows failed:%v\n", err)
		return
	}
	fmt.Printf("n:%v\n", n)
}

func deleteRow(n int) {
	sqlStr := `delete from user where id=?;`
	ret, err := db.Exec(sqlStr, n)
	if err != nil {
		fmt.Printf("delete failed:%v\n", err)
		return
	}
	rowNum, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get row failed:%v\n", err)
		return
	}
	fmt.Printf("delete:%v\n", rowNum)
}

func prepareInsert() {
	sqlStr := `insert into user(name,age) values(?,?);`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	var m = map[string]int{
		"罚款":  789,
		"阿斯顿": 8,
		"爱上了": 89,
		"阿萨德": 39,
	}
	for k, v := range m {
		stmt.Exec(k, v)
		// ret, err := stmt.Exec(k, v)
		// if err != nil {
		// 	fmt.Printf("失败：%v\n", err)
		// 	return
		// }
		// rowNum, err := ret.RowsAffected()
		// if err != nil {
		// 	fmt.Printf("failed:%v\n", err)
		// 	return
		// }
		// fmt.Printf("insert:%v\n", rowNum)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed: %v\n", err)
	}
	fmt.Println("数据库连接成功！")
	// queryOne(1)
	// queryMore(1)
	// insert()
	// updateRow(68, 2)
	// deleteRow(1)
	prepareInsert()
}
