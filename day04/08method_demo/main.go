package main

import "fmt"

// 自定义类型加方法
// 只能给自己包里定义的类型添加方法，不能给别的包里面的类型添加方法
type myInt int

func (i myInt) hello() {
	fmt.Println("int")
}

func main() {
	m := myInt(100)
	m.hello()
}
