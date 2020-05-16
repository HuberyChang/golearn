package main

import "fmt"

// 关闭通道

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	<-ch1
	<-ch1
	x, ok := <-ch1 // 前面已经把通道的数据读完，再读通道数据就是flase，值是chan对应类型的“0”值，int-0; bool-false; string-""
	fmt.Println(x, ok)
	// x, ok = <-ch1
	// fmt.Println(x, ok)
	// x, ok = <-ch1
	// fmt.Println(x, ok)
}
