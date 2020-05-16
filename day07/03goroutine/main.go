package main

import (
	"fmt"
	"time"
)

// goroutine
// 将要并发执行的任务包装成一个函数，调用函数的时候前面加go关键字，就能开启一个goroutine去执行该函数的任务

/*
	goroutine调度模型：
		GMP
			m:n---将m个goroutine跑在n个os线程上
*/
func hello(i int) {
	fmt.Println("kdshfjkas", i)
}

// 程序启动以后会创建goroutine去执行main函数
// main函数结束后，由main启动的goroutine也都结束
// goroutine对应的函数结束了，goroutine结束了
func main() {
	for i := 1; i <= 100; i++ {
		// go hello(i) // 开启一个单独的goroutine去执行hello函数（任务）

		// go func() {
		// 	fmt.Println(i)
		// }()

		go func(i int) {
			fmt.Println(i) // 用的是函数内部的i不是外部的i
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)

}
