package main

import "fmt"

// 空接口，没必要取名字

// interface: 关键字
// interface{}：空接口类型

// 空接口作为函数参数类型
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	// 空接口作为map的值
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)

	m1["name"] = "ahsf家"
	m1["age"] = 249
	m1["merried"] = true
	m1["hobby"] = [...]string{"chang", "tiao", "rap"}
	fmt.Println(m1)

	show(m1)
	show(nil)
	show(false)
}
