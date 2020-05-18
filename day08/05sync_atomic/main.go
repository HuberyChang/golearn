package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	// lock.Lock()
	// x++
	// lock.Unlock()
	atomic.AddInt64(&x, 1)
}

func main() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
