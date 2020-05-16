package main

import "fmt"

//结构体

//声明一个结构体
type people struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个people类型的变量ha
	var ha people
	//通过字段赋值
	ha.name = "chen"
	ha.age = 120
	ha.gender = "哈"
	ha.hobby = []string{"篮球", "离开家啊哈", "阿克纠纷"}
	// 访问变量ha的字段
	fmt.Println(ha)

	// 匿名结构体。多用于临时场景
	var s struct {
		name   string
		age    int
		gender string
		hobby  string
	}
	s.name = "hakj"
	s.age = 9
	s.gender = "难"
	fmt.Println(s)
}
