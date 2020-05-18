package main

import (
	"fmt"
	"sync"
)

// 锁

var x = 0 // 公共资源
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000000; i++ {
		lock.Lock() // 对公共资源加锁
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
