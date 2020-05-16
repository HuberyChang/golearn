package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json
// 1. 序列化:   go语言中的结构体变量--->json格式的字符串
// 2. 反序列化: json格式的字符串--->go语言中的结构体变量

type person struct {
	Name string `json:"name" db:"fjqa" ini:"foaj" ...`
	Age  int    `json:"age"`
	// name string
	// age  int
}

func main() {
	p1 := person{
		Name: "jdqa",
		Age:  18,
		// name: "halh",
		// age:  19,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("failed:%v\n", err)
		return
	}
	fmt.Println(p1)
	fmt.Printf("%#v\n", string(b))

	// 反序列化
	str := `{"name":"asdjfk","age":80}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了在json.Umarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)
}
