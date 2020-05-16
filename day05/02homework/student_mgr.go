package main

import "fmt"

// 学生管理系统
// 有一个物件：
// 	1. 它保存了一些数据	--->结构体的字段
// 	2. 它有一些功能		--->结构体的方法

type student struct {
	id   int64
	name string
}

// 造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentMgr) showStudent() {
	// 从s.allStudent这个map中把所有的学生逐个拎出来
	for _, stu := range s.allStudent { // stu为具体学生
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	// 1. 根据用户输入的内容，创建一个新的学生
	var (
		stuID   int64
		stuName string
	)
	// 获取用户输入
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	// 2. 把新的学生放入s.allStudent这个map中
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功")

}

// 修改学生
func (s studentMgr) editStudent() {
	// 获取用户输入学号
	var stuID int64
	fmt.Printf("请输入学号：")
	fmt.Scanln(&stuID)
	// 展示学号对应的学生信息，没有，查无此人
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人！")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d 姓名：%s\n", stuObj.id, stuObj.name)
	// 请输入修改后的学生名
	fmt.Print("请输入学生的新名字：")
	var newName string
	fmt.Scanln(&newName)
	// 更新学生姓名
	// s.allStudent[stuId].name=newName
	stuObj.name = newName
	s.allStudent[stuID] = stuObj //更新map中的学生
}

// 删除学生
func (s studentMgr) delStudent() {
	// 1. 请用户输入要删除的学生的ID
	var stuID int64
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scanln(&stuID)
	// 2. 去map中找有没有这个学生，没有，查无此人
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("没找到这个人！")
		return
	}
	delete(s.allStudent, stuID)
	fmt.Println("你牛逼！")
}
