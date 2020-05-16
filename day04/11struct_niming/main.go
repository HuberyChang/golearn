package main

import "fmt"

// 匿名字段
// 不常用！！！

type person struct {
	string
	int
}

func main() {
	p1 := person{
		"klaj",
		18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string) //把类型当做字段
}
