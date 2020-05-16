package main

import "fmt"

// 闭包？
// 闭包首先他是一个函数，其次它包含一个他外部作用域的一个变量

// 底层原理：
// 1.函数可以作为返回值
// 2.函数内部查找变量的顺序，先在内部查找，再到外部

func adder() func(int) int { //基于return的返回值类型，adder的返回值也是一个带一个int类型参数和一个int类型返回值的函数
	x := 100
	return func(y int) int { // return返回值是一个匿名函数，带一个int类型参数和int类型返回值
		x += y //x不在函数内部，需要到外部查找
		return x
	}
}

func adder1(x int) func(int) int { //基于return的返回值类型，adder的返回值也是一个带一个int类型参数和一个int类型返回值的函数
	return func(y int) int { // return返回值是一个匿名函数，带一个int类型参数和int类型返回值
		x += y
		return x
	}
}

func adder2(x int) func() int { //基于return的返回值类型，adder的返回值也是一个带一个int类型返回值的函数
	// x := 100
	return func() int { // return返回值是一个匿名函数，带一个int类型返回值
		y := 200
		x += y
		return x
	}
}

func adder3() func() int { //基于return的返回值类型，adder的返回值也是一个带一个int类型返回值的函数
	x := 100
	return func() int { // return返回值是一个匿名函数，带一个int类型返回值
		y := 200
		x += y
		return x
	}
}

func main() {

	fmt.Println("========adder=========")
	ret := adder() // adder()函数执行完返回值是一个函数，用ret来接收，所以ret是一个带一个int类型参数和一个int类型返回值的函数
	// ret := adder(100)
	retn := ret(200)  //给ret一个int类型的参数，返回一个int类型的返回值
	fmt.Println(retn) //打印这个返回值
	fmt.Println("========adder=========")

	fmt.Println("========adder1=========")
	ret1 := adder1(100) //ret1不仅包含返回值函数，还包含一个外部变量x，因为x并不是return的函数的内部变量
	ret21 := ret1(200)
	fmt.Println(ret21)
	fmt.Println("========adder1=========")

	fmt.Println("========adder2=========")
	ret2 := adder2(100)
	ret22 := ret2()
	fmt.Println(ret22)
	fmt.Println("========adder2=========")

	fmt.Println("========adder3=========")
	ret3 := adder3()
	ret32 := ret3()
	fmt.Println(ret32)
	fmt.Println("========adder3=========")

}
