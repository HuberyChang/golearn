package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 日志demo

func main() {
	obj, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	log.SetOutput(obj)
	for {
		log.Println("log test...")
		time.Sleep(time.Second * 3)
	}

}
