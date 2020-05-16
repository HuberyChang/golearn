package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看、新增、删除学生
*/

type student struct {
	id   int64
	name string
}

var (
	allStudents map[int64]*student //声明变量
)

// newStudent是student的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudents() {
	// 遍历学生打印
	for k, v := range allStudents {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func addStudent() {
	// 向allStudents中添加一个新学生
	// 1.创建一个新学生
	// 1.1 获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Print("学号：")
	fmt.Scanln(&id)
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	// 1.2 造学生（调用student的构造函数）
	newStu := newStudent(id, name)
	// 2. 追加到allStudents这个map中
	allStudents[id] = newStu
}

func delStudent() {
	var (
		deleteID int64
	)
	fmt.Print("学号：")
	fmt.Scanln(&deleteID)
	// 2.去allStudents这个map中根据学号删除对应的键值对
	delete(allStudents, deleteID)
}

func main() {
	allStudents = make(map[int64]*student, 48) //初始化，开辟内存空间
	// 1.打印菜单
	for {
		fmt.Println("欢迎光临学生管理系统")
		fmt.Println(`
		1.查看
		2.新增
		3.删除
		4.退出
		`)
		fmt.Print("请输入：")
		// 2.等待用户选择要干嘛
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d的是\n", choice)
		// 3.执行对应操作
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚！")
		}
	}
}
