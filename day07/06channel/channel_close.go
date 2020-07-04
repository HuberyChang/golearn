package main

import (
	"fmt"
	"strconv"
)

func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //通知对方，通道关闭，必须的，否则会死锁！！！
}

func sendData1(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子goroutine的第 %d 个数据\n", i)
	}
	close(ch)
}

func main() {
	//var ch1 chan int
	//ch1 = make(chan int)
	//go sendData(ch1)

	//读取通道数据

	//1.for循环访问同通道
	//for {
	//	time.Sleep(time.Second)
	//	v, ok := <-ch1
	//	if !ok {
	//		fmt.Println("通道数据取完了。。。", ok)
	//		break
	//	}
	//	fmt.Println("读取的数据", v, ok)
	//}
	//fmt.Println("main ... over...")

	//2.range循环访问通道
	//for v := range ch1 {
	//	time.Sleep(time.Second)
	//	fmt.Println("读取的数据为：", v)
	//}
	//fmt.Println("main...over...")

	ch3 := make(chan string, 4)
	go sendData1(ch3)
	fmt.Println("++++++++++", len(ch3))

	for {
		v, ok := <-ch3
		if !ok {
			fmt.Println("数据取玩了", ok)
			break
		}
		fmt.Println("========", len(ch3))
		fmt.Println("读取的数据位：", v, ok)
	}
	fmt.Println("main...over++++")

}
