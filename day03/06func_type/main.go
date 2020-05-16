package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("jdskf")
}

func f2() int {
	return 10
}

// 函数也可作为参数类型
func f3(x func() int) { //参数x是func()int类型的
	ret := x()
	fmt.Println(ret)
}

func f4(x, y int) int {
	return x + y
}

func ff(a, b int) int {
	return a + b
}

//函数还可以是返回值
func f5(x func() int) func(int, int) int { //接收一个func()int类型的参数，返回值是func(int,int)int
	// ret := func(a, b int) int {
	// 	return a + b
	// }
	// return ret
	return ff
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)

	f3(f2) //参数f2符合函数f3的定义
	f3(b)  //参数b符合函数f3的定义

	f7 := f5(f2)
	fmt.Printf("%T", f7)

}
