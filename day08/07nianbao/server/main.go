package main

import (
	"bufio"
	"fmt"
	prt "golearn/day08/07nianbao/solvenianbao"
	"io"
	"net"
)

func processor(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// var buf [1024]byte
	for {
		// n, err := reader.Read(buf[:])
		// if err == io.EOF {
		// 	break
		// }
		// if err != nil {
		// 	fmt.Println("read from client failed,", err)
		// 	return
		// }
		// recvStr := string(buf[:n])
		recvStr, err := prt.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read from client failed,", err)
			return
		}
		fmt.Println("receive client:", recvStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("lsiten failed", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
			continue
		}
		go processor(conn)
	}
}
