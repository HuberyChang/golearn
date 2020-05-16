package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("f1 func")
	f()
}

func f2(x, y int) {
	fmt.Println("f2 func")
	fmt.Println(x + y)
}

// 定义一个函数对f2包装
func f3(f func(int, int)) func() { // f3的返回值类型基于return返回的返回值类型，因为tmp是一个没有参数和返回值的函数，所以f3的返回值也要是一个没有参数和返回值的函数
	tmp := func() {
		// fmt.Println("hello")
		f(100, 200)
	}
	return tmp //tmp是匿名函数，作为返回值，没有参数没有返回值的函数
	// return tmp() //tmp后面加了（）相当于变成执行tmp函数了
}

func f4(f func(int, int), x, y int) func() { // f4的返回值类型基于return返回的返回值类型，因为tmp是一个没有参数和返回值的函数，所以f4的返回值也要是一个没有参数和返回值的函数
	tmp := func() {
		// fmt.Println("hello")
		f(x, y)
	}
	return tmp //tmp是匿名函数，作为返回值，没有参数没有返回值的函数
	// return tmp() //tmp后面加了（）相当于变成执行tmp函数了
}

func main() {
	// f1(f2) //f2不能作为f1的参数，因为与f1的参数类型不匹配
	f1(f3(f2))
	fmt.Println("++++++++++++++++++f4+++++++++")
	ret := f4(f2, 100, 200)
	f1(ret)
	fmt.Println("++++++++++++++++++f4+++++++++")
}
