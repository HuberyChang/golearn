package main

import "fmt"

// 结构体嵌套
// 结构体的匿名变量--->结构体的嵌套--->结构体的匿名嵌套--->匿名嵌套结构体的字段冲突
type address struct {
	dizhi string
	city  string
}

type person struct {
	name    string
	age     int
	address //匿名嵌套结构体
	workplace
	// dizhi string
	// city string
}

type company struct {
	name string
	addr address //结构体嵌套
	// dizhi string
	// city string
}

type workplace struct {
	dizhi string
	city  string
}

func main() {
	p1 := person{
		name: "hdkaf",
		age:  19,
		address: address{
			dizhi: "sdj7013",
			city:  "9032jjds",
		},
		workplace: workplace{
			dizhi: "淦",
			city:  "抓呢",
		},
	}
	fmt.Println(p1)
	/*fmt.Println(p1.city)  匿名嵌套函数直接调用，先在自己结构体内查找这个字段，找不到再去匿名的结构体中查找该字段
	字段冲突 */
	fmt.Println(p1.workplace.city, p1.address.city)

	ps := company{
		name: "tanjie",
		addr: address{
			dizhi: "fengde",
			city:  "chengdu",
		},
	}
	fmt.Println(ps)
	fmt.Println(ps.name, ps.addr.dizhi)
}
