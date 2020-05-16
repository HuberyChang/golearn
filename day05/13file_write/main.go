package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容

func write1() {

	obj, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open failed,err:%v", err)
		return
	}
	obj.Write([]byte("fasdhasf\n"))
	obj.WriteString("就安理会代发\n")
	obj.Close()
}

// bufio写文件
func write2() {
	obj, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open failed,err:%v", err)
		return
	}
	defer obj.Close()
	// 创建一个写对象
	wr := bufio.NewWriter(obj)
	wr.WriteString("a9a34ufp9a8") // 写到缓存
	wr.Flush()
}

// ioutil写文件
func write3() {
	str := "安利好伐啦"
	err := ioutil.WriteFile(".test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}

func main() {
	// write1()
	// write2()
	write3()
}
