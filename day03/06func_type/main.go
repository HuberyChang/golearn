package main

import "fmt"

/*
	函数类型
	高阶函数：
		go语言允许将一个函数作为另外一个函数的参数

	func1(),func2()
	将func1作为函数func2的参数
		func2函数：就是高阶函数
			接收一个函数作为参数的函数，就是高阶函数
		func1函数：就是回调函数
			作为另外一个函数的参数的函数，就是回调函数
*/

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

//加法运算
func add(a, b int) int {
	return a + b
}

//减法运算
func sub(a, b int) int {
	return a - b
}

func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	ret := fun(a, b)
	return ret
}

func main() {
	a := f1
	fmt.Printf("a===%T\n", a)
	b := f2
	fmt.Printf("b===%T\n", b)

	f3(f2) //参数f2符合函数f3的定义
	f3(b)  //参数b符合函数f3的定义

	f7 := f5(f2)
	fmt.Printf("===++%T\n", f7)

	fmt.Println("================")
	fmt.Printf("add===%T\n", add)
	fmt.Printf("sub===%T\n", sub)
	fmt.Printf("oper===%T\n", oper)

	ret1 := add(10, 20)
	fmt.Println(ret1)
	ret2 := oper(10, 20, add)
	fmt.Println(ret2)
	ret3 := oper(20, 10, sub)
	fmt.Println(ret3)
	//乘法，匿名函数
	fun1 := func(a, b int) int {
		return a * b
	}
	ret4 := oper(1, 3, fun1)
	fmt.Println(ret4)
	//除法，匿名函数
	ret5 := oper(100, 20, func(a int, b int) int {
		if b == 0 {
			fmt.Println("除数为0")
			return 0
		}
		return a / b
	})
	fmt.Println(ret5)
	fmt.Println("================")

}
