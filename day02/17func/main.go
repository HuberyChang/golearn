package main

import "fmt"

// 函数

// 函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func f1(x int, y int)  {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2()  {
	fmt.Println("f2")
}

// 没有参数
func f3() (ret int)  {
	return 3
}

// 返回值可以命名也可以不命名
// 命名返回值就相当于在函数中声明一个变量
func f4(x int, y int)(ret int)  {
	ret = x + y
	return 	//使用命名返回值可以return后省略
}

func f5() (int)  {
	ret := 3
	return ret	//未使用命名返回值，return后不可省略
}


// 多个返回值
func f6()(int, string)  {
	return 1, "haha"
}

// 参数类型简写，当参数中连续多个参数类型一致，可以将前面些参数类型省略
func f7(x, y int, m, n string, i, j bool)(int)  {
	return x + y
}

// 可变长参数
func f8(x string, y ...int)  {
	fmt.Println(x)
	fmt.Println(y)
}
func main() {
	r := sum(1, 10)
	f1(2, 3)
	f2()
	f3()
	// f4()
	fmt.Println(r)
	_, n := f6()
	fmt.Println(n)
	f8("xiayule", 1, 3, 4, 5)
}
