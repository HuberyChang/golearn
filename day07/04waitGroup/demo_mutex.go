package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10
var wg1 sync.WaitGroup
var mutex sync.Mutex // 创建锁头

//var mutex = new(sync.Mutex)

func saleTicket(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg1.Done()
	for {
		//上锁
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Println(name, "售出:", ticket)
			ticket--
		} else {
			mutex.Unlock() // 条件不满足，也要解锁，让别的goroutine可以执行
			fmt.Println(name, "售罄，没票。。。")
			break
		}
		mutex.Unlock() //解锁
	}
}

func main() {
	wg1.Add(4)

	//模拟4个窗口
	go saleTicket("1:")
	go saleTicket("2:")
	go saleTicket("3:")
	go saleTicket("4:")

	wg1.Wait() //main要等待
	//time.Sleep(time.Second)

}
