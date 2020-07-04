package main

import (
	"fmt"
)

//var wg sync.WaitGroup // 创建同步等待组的对象

func fun1() {
	//方式一，在代码前：defer wg.Done()
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("fun1()函数中打印。。A", i)
	}
}

func fun2() {
	for i := 0; i < 10; i++ {
		fmt.Println("====fun2()函数中打印。。B", i)
	}
	//方式二，在代码尾
	wg.Done() // 给wg等待组中的counter数值减1
}

func main() {
	/*
		WaitGroup:同步等待组
		Add():设置等待组中要执行的子goroutine的数量
		Wait():让主goroutine处于等待
		Done():让等待组中的counter计数器的值减1，同Add(-1)
	*/
	wg.Add(2)
	go fun1()
	go fun2()
	fmt.Println("main进入阻塞状态。。。等待wg中的子goroutine结束")
	wg.Wait()
	fmt.Println("main。。解除阻塞")
}
