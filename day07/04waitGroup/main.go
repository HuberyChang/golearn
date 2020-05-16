package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

func f() {
	rand.Seed(time.Now().UnixNano()) // 保证每次执行的结果不一样
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		// rand.Intn(10) // 0<= x <10
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
	fmt.Println(i)
	// wg.Done()
}

var wg sync.WaitGroup

func main() {
	// wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1) // 逐次加1，直到为10
		go f1(i)
	}
	// 如何知道这10个goroutine都结束
	wg.Wait() // 等待wg的计数器减为0
}
