package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock

var (
	x      = 0
	lock   sync.Mutex
	rwlock sync.RWMutex
	wg     sync.WaitGroup
)

// 读操作，加读锁
func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	rwlock.RLock()
	// rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	rwlock.RUnlock()
	// rwlock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
