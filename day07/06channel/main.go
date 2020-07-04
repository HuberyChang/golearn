package main

import (
	"fmt"
	"sync"
)

var b chan int // 需要制定通道中元素的类型
var wg sync.WaitGroup

func noBufChan() {
	fmt.Println(b)
	//创建通道
	b = make(chan int)
	wg.Add(1)

	//启动一个子goroutine
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("取到了x的值。。。", x)
	}()

	//向通道写数据
	b <- 10
	fmt.Println("发送到chan中。。。")

	wg.Wait()
}

func haveBufChan() {
	fmt.Println(b)
	b = make(chan int, 1) // 容量为1
	b <- 10
	fmt.Println("10 sent to chan...")
	// b <- 20
	// fmt.Println("20 sent to chan...")
	x := <-b
	fmt.Println("从通道b中取到", x)
	close(b)
}

func test(ch chan int) {
	fmt.Printf("%T,%v\n", ch, ch)
}

func main() {
	// noBufChan()
	haveBufChan()
	var a = make(chan int)
	fmt.Println(a) //0xc000086060
	test(a)        //chan int,0xc000086060
}
