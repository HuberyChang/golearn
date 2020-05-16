package main

import "fmt"


// go语言推荐驼峰命名
// var student_name string  //下划线格式
// var studentName string	 //小驼峰
// var StudentName string	 //大驼峰


// 单独声明变量
// var name string
// var age int
// var isOk bool

// 批量声明全局变量
var (
	name string  // ""
	age int      // 0
	isOk bool    // false
)

// go语言中声明非全局变量必须使用，否则报错
func main()  {
	name = "理想"
	age = 6
	isOk = true
	fmt.Print(isOk)   			 // 在终端中输出要打印为内容
	fmt.Println()
	fmt.Printf("name:%s", name)  // %s：占位符，使用name这个变量的值去替换占位符
	fmt.Println()
	fmt.Println(age)   			 // 打印完指定内容会在后面加一个换行符

	// 声明变量同时赋值
	var s1 string = "hehehe"
	fmt.Println(s1)

	// 类型推导
	var s2 = "xcxz"
	fmt.Println(s2)

	//简短变量声明，只能在函数里用
	s3 := "hfquweio23"
	fmt.Println(s3)
}