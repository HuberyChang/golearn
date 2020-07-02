package main

import "fmt"

//defer

func f1() int { //返回值只知道类型，未命名
	x := 5
	defer func() {
		x++ //修改的是x不是返回值
	}()
	return x //1.返回值赋值  2.defer  3.真正的RET指令
}

func f2() (x int) { //返回值既指定类型右指定名称
	defer func() {
		x++ //defer 修改的是返回值x
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ //修改的是x
	}()
	return x //返回值=y=x=5  2.defer修改x的值  3.真正RET指令
}

func f4() (x int) {
	defer func(x int) {
		x++ //改变的是函数中的x的副本
	}(x)
	return 5 //返回值x=5=x
}

func f5() (x int) {
	defer func(x *int) {
		(*x)++ //x现在是*int类型，（*x）=**int，内存地址指向的值+1
	}(&x) //传入x变量地址
	return 5 //返回值x=5=x
}

func f6() (x int) {
	defer func(x int) int {
		x++ //改变的是函数中的x的副本
		return x
	}(x)
	return 5 //返回值x=5=x
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
	fmt.Println(f5()) //6
	fmt.Println(f6()) //6
	fmt.Println("++++++++++++++++++++++")
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
