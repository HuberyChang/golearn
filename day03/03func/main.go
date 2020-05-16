package main

import "fmt"

// 函数

func f1() {
	fmt.Println("hkjlfahld")
}

func f2(name string) {
	fmt.Println("hello", name)
}

// 带参数、返回值
func f3(x int, y int) int {
	sum := x + y
	return sum
}

// 参数类型简写
func f4(x, y int) int {
	return x + y
}

//可变参数
func f5(x string, y ...int) int {
	fmt.Println(y) //y是一个int类型的切片
	return 1
}

//命名返回值
func f6(x, y int) (sum int) {
	sum = x + y // 如果使用命名返回值，那么函数中可以直接使用返回值变量
	return      //return 后可以省略返回值变量
}

//go语言支持多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}


func main() {
	f1()
	f2("h789")
	// fmt.Println(f3(2, 4))
	ret := f3(2, 4)
	fmt.Println(ret)

	f5("hfajkl", 1, 2, 3, 4)
	// func f8()  {//不能在命名函数中再声明命名函数
		
	// }
	}