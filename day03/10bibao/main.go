package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("f1")
	f()
}

func f2(x, y int) {
	fmt.Println("f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func main() {
	ret := f3(f2, 100, 200)
	f1(ret)
}
