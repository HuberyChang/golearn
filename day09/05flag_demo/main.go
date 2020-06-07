package main

import (
	"flag"
	"fmt"
	"time"
)

// flag获取参数命令行
func main() {
	name := flag.String("name", "hahah", "请输入名字：")
	age := flag.String("age", "19", "请输入年龄：")
	married := flag.Bool("married", false, "婚否：")
	cTime := flag.Duration("time", time.Second, "结婚年龄：")
	flag.Parse() // 先解析再使用
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)
	fmt.Println(flag.Args())  // 返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  // 返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的命令行参数个数

	// 注释
	// var name string
	// flag.StringVar(&name, "name", "我", "请输入名字：")
	// flag.Parse()
	// fmt.Println(name)
}
