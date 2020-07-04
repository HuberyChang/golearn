package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg2 *sync.WaitGroup

func readData(i int) {
	defer wg2.Done()
	fmt.Println(i, "开始：read start")
	rwMutex.RLock()
	fmt.Println(i, "正在读取数据：reading。。。")
	time.Sleep(3 * time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println(i, "读操作结束：read over。。。")
}

func writeData(i int) {
	defer wg2.Done()
	fmt.Println(i, "开始：write start")
	rwMutex.Lock()
	fmt.Println(i, "正在写数据：writing。。。")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock() //读操作解锁
	fmt.Println(i, "写操作结束：write over。。。")
}

func main() {
	rwMutex = new(sync.RWMutex)
	wg2 = new(sync.WaitGroup)

	wg2.Add(3)
	go readData(1)
	go readData(2)
	go writeData(3)
	wg2.Wait()
	fmt.Println("main over")
}
