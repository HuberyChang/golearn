package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么要用context
var wg sync.WaitGroup

// var notify bool
var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("fjal")
		time.Sleep(time.Millisecond * 500)
		// if notify{
		// 	break
		// }

		select {
		case <-exitChan:
			break FORLOOP
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// 通知子goroutine退出
	exitChan <- true
	// notify=true
	wg.Wait()
}
