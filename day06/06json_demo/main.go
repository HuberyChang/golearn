package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"` //反射到Name
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"adshf","age":9}` //	此处的name对应结构体的name，结构体的name又对应到Name字段（反射）
	var p person
	json.Unmarshal([]byte(str), &p) //传进一个指针。可以动态的去结构体查看有哪些字段
	fmt.Println(p.Name, p.Age)
}
