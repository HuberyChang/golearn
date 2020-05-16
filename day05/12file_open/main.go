package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件

// 循环读取
func readFromFileByxuanhuan() {
	obj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!", err)
		return
	}
	defer obj.Close()
	// var tmp = make([]byte, 128)	// 指定读的长度
	var tmp [128]byte //简便定义
	for {
		n, err := obj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Printf("err:%v", err)
			return
		}
		fmt.Printf("%d个字节\n", n)
		fmt.Println(string(tmp[:]))
		if n == 0 {
			return
		}
	}
}

// 利用bufio这个包读取文件
func readFromFilebyBufio() {
	obj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	// 记得关闭文件
	defer obj.Close()
	// 创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(obj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("err:%v", err)
			return
		}
		fmt.Print(line)
	}
}

// 利用ioutil读文件
func readFromFileByIoutil() {
	obj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	defer obj.Close()
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// readFromFileByxuanhuan
	// readFromFilebyBufio()
	readFromFileByIoutil()
}
