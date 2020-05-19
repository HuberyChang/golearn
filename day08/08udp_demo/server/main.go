package main

import (
	"fmt"
	"net"
	"strings"
)

// udp

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 4000,
	})
	// defer conn.Close() 如果出错就关不了
	if err != nil {
		fmt.Println("udp failed,err", err)
		return
	}
	defer conn.Close() // 先处理错误，没错误才关闭
	// 不需要建立连接
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:]) // data[:]是把数组data转化为切片
		if err != nil {
			fmt.Println("read from udp failed", err)
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		// 发送数据
		conn.WriteToUDP([]byte(reply), addr)
	}
}
