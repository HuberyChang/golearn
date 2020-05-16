package main

import "fmt"

//defer，多用于函数结束前释放资源（文件句柄，数据库连接，socket连接）

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
	}()
	return x
}

func f2() (x int) { //返回值=x，返回值赋值，x=5，再执行defer，再执行返回真正的返回值
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) { //返回值=y，返回值赋值，y=x=5，再执行defer，但是y已经赋值为5，此时x=6
	x := 5
	defer func() {
		x++ //修改的是x
	}()
	return x
}

func f4() (x int) { //返回值=x，返回值赋值，x=5，再执行defer，再执行返回真正的返回值
	defer func(x int) { //函数中的x跟外层函数的x不是同一个
		x++ //改变的是函数的副本
	}(x)
	return 5 //返回值=x=5
}

func main() {
	// deferDemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
