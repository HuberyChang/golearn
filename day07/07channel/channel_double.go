package main

import "fmt"

func sendData(ch chan string, done chan bool) {
	ch <- "合法" //发送

	data := <-ch //读取
	fmt.Println("主", data)
	done <- true
}

func main() {
	ch1 := make(chan string)
	done := make(chan bool)
	go sendData(ch1, done)

	data := <-ch1 //读取
	fmt.Println("子", data)

	ch1 <- "家得福" //发送
	<-done
}
