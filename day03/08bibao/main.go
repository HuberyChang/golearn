package main

import (
	"fmt"
)

/*
	闭包
	go语言支持函数式编程：
		支持将一个函数作为另一个函数的参数
		也支持将一个函数作为另一个函数的返回值
	闭包：
		一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量（外层函数中的参数，或者外层函数中直接定义的变量），并且该外层函数的返回值就是该内层函数

		这个内层函数和外层函数的局部变量，统称闭包结构

		局部变量的生命周期会发生变化，正常的局部变量随函数调用而创建，随函数结束而销毁。
		但是闭包结构中的外层函数的局部变量并不会随外层函数的结束而销毁因为内层函数还会继续使用。
*/

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
	fmt.Println("地址：", tmp) //地址： 0x4a0400
	return tmp              //tmp是匿名函数，作为返回值，没有参数没有返回值的函数
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

func increment() func() int {
	//1.定义一个局部变量
	i := 0
	//2.定义一个匿名函数，给变量自增并返回
	fun := func() int {
		i++
		return i
	}
	//3.返回该匿名函数
	return fun
}

func main() {
	// f1(f2) //f2不能作为f1的参数，因为与f1的参数类型不匹配
	f1(f3(f2))
	fmt.Println("++++++++++++++++++f4+++++++++")
	ret := f4(f2, 100, 200)
	f1(ret)
	fmt.Println("++++++++++++++++++f4+++++++++")
	ret1 := increment()      // ret=fun
	fmt.Printf("%T\n", ret1) //func() int
	fmt.Println(ret1)        // 0x4a0340
	v1 := ret1()
	fmt.Println(v1) // 1
	v2 := ret1()
	fmt.Println(v2) // 2
	fmt.Println(ret1())
	fmt.Println(ret1())
	fmt.Println(ret1())
	fmt.Println(ret1())
	ret2 := increment()
	fmt.Println(ret2) // 0x4a0420
	v3 := ret2()
	fmt.Println(v3) // 1

}
