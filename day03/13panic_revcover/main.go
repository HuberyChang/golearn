package main

import "fmt"

// panic & recover

func funcA() {
	fmt.Println("a")
}

func funcB() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接...")
	}()
	panic("error!") //程序崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
