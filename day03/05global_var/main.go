package main

import "fmt"

// 作用域

var x = 100 //定义一个全局变量

//定义一个函数
func f1() {
	name := "haha" //内部定义变量只能在内部使用
	// 函数中查找变量的顺序
	// 1.先在函数内部查找，找不到则往函数外找
	fmt.Println(x, name)
}

func main() {
	f1()
	if i := 10; i < 18 { //i只能在if循环内
		fmt.Println("未成年")
	}
}
