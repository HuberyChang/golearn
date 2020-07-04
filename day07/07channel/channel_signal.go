package main

import (
	"fmt"
)

func main() {
	//ch1 := make(chan int) //双向，读，写
	ch2 := make(chan<- int) //单向，只写
	//ch3 := make(<-chan int) //单向，只读

	//ch2 <- 100
	//data := <-ch2 //invalid operation: <-ch2 (receive from send-only type chan<- int)
	//ch3 <- 1000 //invalid operation: ch3 <- 1000 (send to receive-only type <-chan int)

	//go fun1(ch1)
	go fun1(ch2)
	//data := <-ch2
	//fmt.Println(data)
}

//该函数，只能操作只写通道
func fun1(ch chan<- int) {
	ch <- 100
	fmt.Println("fun1函数结束。。。")
}
