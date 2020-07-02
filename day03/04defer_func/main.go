package main

import (
	"fmt"
)

/*
	1.defer，多用于函数结束前释放资源（文件句柄，数据库连接，socket连接）
	2.defer函数传递参数的时候，defer函数调用的时候已经传递参数，只是暂时不执行函数中的代码（defer是延迟函数执行，不是延迟函数调用）
	注意：
		当外围函数中的语句正常执行完毕，只有其中所有的defer函数都执行完毕，外围函数才会真正的结束执行。
		当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回
		当外围函数中的代码发生panic时，只有其中所有的延迟函数都执行完毕后，该panic才会真正被扩展至调用函数
*/

func deferDemo() {
	fmt.Println("1")
	defer fmt.Println("11") //defer将他后面的语句延迟到函数即将返回的时候执行，多个defer之间语句逆序执行
	defer fmt.Println("12") //defer将他后面的语句延迟到函数即将返回的时候执行
	defer fmt.Println("13") //defer将他后面的语句延迟到函数即将返回的时候执行
	fmt.Println("2")
}

// go语言中函数的return不是原子操作，在底层是分为两步执行
// 第一步：返回值赋值
// defer
// 第二步：真正的ret返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func f1() int { //未命名函数，直接 返回值赋值
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
		fmt.Println("f1:", x)
	}()
	return x
}

func f2() (x int) { //返回值=x，返回值赋值，x=5，再执行defer，再执行返回真正的返回值
	defer func() {
		x++
		fmt.Println("f2:", x)
	}()
	return 5
}

func f3() (y int) { //返回值=y，返回值赋值，y=x=5，再执行defer，但是y已经赋值为5，此时x=6
	x := 5
	defer func() {
		x++ //修改的是x
		fmt.Println("f3:", x)
	}()
	return x
}

func f4() (x int) { //返回值=x，返回值赋值，x=5，再执行defer，再执行返回真正的返回值
	defer func(x int) { //函数中的x跟外层函数的x不是同一个
		x++ //改变的是函数的副本
		fmt.Println("f4:", x)
	}(x)
	return 5 //返回值=x=5
}

//defer函数传递参数的时候，defer函数调用的时候已经传递参数，只是暂时不执行函数中的代码
func f5(a int) {
	fmt.Println("f5()函数中打印a：", a) //2
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	// deferDemo()
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	a := 2
	fmt.Println("1.main中打印a:===", a) //2
	defer f5(a)
	a++
	fmt.Println("2.main中打印a:===", a) //3
	fmt.Println("=============")
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
