package main

import (
	"fmt"
	"io"
	"os"
)

func f1() {
	// 打开要操作的文件
	obj, err := os.OpenFile("./test.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	// defer obj.Close()

	// 因为没办法直接在原始文件中加插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./test.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	defer tmpFile.Close()

	// 读取源文件写入临时文件
	var ret [1]byte
	n, err := obj.Read(ret[:])
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	// 写入临时文件
	tmpFile.Write(ret[:n])

	// 再写入要插入的内容
	var s []byte
	s = []byte{'9'}
	tmpFile.Write(s)

	// 紧接着把原文件后续的也写入了临时文件中
	var x [1024]byte
	for {
		n, err := obj.Read(x[:n])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("err:%v\n", err)
			return
		}
		tmpFile.Write(x[:n])
	}

	// 原文件后续的也写入了临时文件中
	obj.Close()
	tmpFile.Close()
	os.Rename("./test.tmp", "./test.txt")

	// fmt.Println(string(ret[:n]))
	// obj.Seek(2, 0) //相对文件头，光标移到一个位置
	// var ret [i]byte
	// n, err :=

}

func main() {
	f1()
}
