package main

import (
	"fmt"
	"os"
)

var smr studentMgr // 声明一个全局变量管理学生对象

func showMenu() {
	fmt.Println("+++++++++++++++++welcome sms!++++++++++++++++")
	fmt.Println(`
	1.查看学生
	2.增加学生
	3.修改学生
	4.删除学生
	5.退出
	`)
}

func main() {
	smr = studentMgr{ //修改的全局的那个变量
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		// 等待输入
		fmt.Print("请输入要执行的序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输的是：", choice)
		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.delStudent()
		case 5:
			os.Exit(1)
		default:
			break
		}
	}
}
