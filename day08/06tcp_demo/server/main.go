package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server端

func processor(conn net.Conn) {
	// 与客户端建立连接
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed,err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		fmt.Printf("回话：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}

func main() {
	// 1. 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("tcp failed,err:", err)
		return
	}
	// 2. 等待别人来连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn failed err:", err)
			return
		}
		// 3. 与客户端通信
		go processor(conn)
	}
}
