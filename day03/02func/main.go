package main

import "fmt"

// 函数

// 函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数
func f3() (ret int) {
	return 3
}

/*
	return
	语义：
		1.有返回值，就要使用return将返回值返回调用处
		2.同时意味着结束函数执行
	注意：
		1.返回值可以命名也可以不命名
		2.命名返回值就相当于在函数中声明一个变量
		3.函数有返回值，必须用return将结果返回给调用处，return后面的数据必须和定义的一致：个数、类型、顺序
		4.如果一个函数定义了返回值，那么函数中有分支，需要保证：无论执行哪个分支都要有return语句执行到
		5.如果一个函数没定义返回值，也可使用return来结束函数
*/
func f4(x int, y int) (ret int) {
	ret = x + y
	return //使用命名返回值可以return后省略
}

func f5() int {
	ret := 3
	return ret //未使用命名返回值，return后不可省略
}

// 多个返回值
func f6() (int, string) {
	return 1, "haha"
}

// 参数类型简写，当参数中连续多个参数类型一致，可以将前面些参数类型省略
func f7(x, y int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数
func f8(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}
func f9() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			//break  // 跳出循环，会执行 fmt.Println("f9.....")
			return // 在哪执行return语句，函数在哪结束，不会执行 fmt.Println("f9.....")
		}
	}
	fmt.Println("f9.....")
}

func main() {
	r := sum(1, 10)
	f1(2, 3)
	f2()
	f3()
	// f4()
	fmt.Println(r)
	_, n := f6() // 需要两个变量接收返回值
	fmt.Println(n)
	f8("xiayule", 1, 3, 4, 5)
	f9()            //调用函数f9()
	fmt.Println(f9) //打印函数f9申请的内存地址
}
