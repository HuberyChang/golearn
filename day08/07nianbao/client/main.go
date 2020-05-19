package main

import (
	"fmt"
	prt "golearn/day08/07nianbao/solvenianbao"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dail failed", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `hello, Hello.How are you?`
		// 调用协议编码数据
		b, _ := prt.Encode(msg)
		conn.Write(b)
		// time.Sleep(time.Second)
	}
}
