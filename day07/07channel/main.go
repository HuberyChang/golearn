package main

import (
	"fmt"
	"sync"
)

// channel 练习
// 1. 启动一个goroutine，生成100个书法送到ch1
// 2. 启动一个goroutine，从ch1中取值，计算平方放到ch2
// 3. 在main中，从ch2取值打印

var wg sync.WaitGroup
var once sync.Once

// 单向通道：限制通道作为参数只能一种操作，要么读要么存
// func f1(ch1 chan<- int)---只能往ch1存值
// func f1(ch1 <-chan int)---只能从ch1取值

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() {
		close(ch2)
	})
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 200)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
