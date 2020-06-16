package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么要用context
var wg sync.WaitGroup

// var notify bool
var exitChan = make(chan bool, 1)

func f(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("fjal")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func main() {
	ctx, cancle := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	// 通知子goroutine退出
	cancle()
	wg.Wait()
}
