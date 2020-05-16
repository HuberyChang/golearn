package main

import "fmt"

func f1(name string) {
	fmt.Println("发斯蒂芬:", name)
}

func f2(f func(string), name string) {
	f1(name)
}

func f3() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func f4(f func()) {
	f()
}

func f5(f func(string), name string) func() {
	return func() {
		f1(name)
	}
}

func main() {
	f2(f1, "kjhfakj")
	ret := f3()
	fmt.Println(ret(1, 2))

	fc := f5(f1, "hahah")
	f4(fc)
}
