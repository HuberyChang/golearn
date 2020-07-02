package main

import "fmt"

/*
	匿名函数：没名字的函数

	定义一个匿名函数，直接在函数后面加"（）"进行调用。通常只能使用一次。但是如果赋值给一个变量，则可多次调用
	一次：
		func(){
			fmt.Println("func")
		}
	多次：
		f1:=func(){	// f1保存的是匿名函数的内存地址
			fmt.Println("f1")
		}
		f1()
		f1()

	匿名函数：
		go语言支持函数式编程：
		1.将匿名函数作为另一个函数的参数，回调函数
		2.将匿名函数作为另一个函数的返回值，可以形成闭包结构

*/

// var f1 = func(x, y int) {
// 	fmt.Println(x + y)
// }

func main() {
	// f1(10, 20)
	// 函数内部没有办法声明带名字的函数，只能有匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	ret1 := func(a, b int) int {
		return a + b
	}(10, 20) // 匿名函数调用了，将执行结果返回给ret1
	fmt.Println("ret1=======:", ret1)

	ret2 := func(a, b int) int {
		return a + b
	} // 将匿名函数的内存地址，赋值给ret2
	fmt.Println("ret2=======:", ret2)
	fmt.Println("ret2值=======:", ret2(100, 200)) // 函数调用

	// 如果只是调用一次的函数，还可以简写成立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("hello world")
	}(100, 200)
}
