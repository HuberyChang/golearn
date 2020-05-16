package main

import (
	"encoding/json"
	"fmt"
)

// 复习结构体

type a struct { /* 有名字结构体，“{}”是结构体本身的花括号 */
	x int
	y string
}

type person struct {
	name string
	age  int
}

// 构造（结构体变量）函数，返回值是对应类型的结构体。构造函数解决结构体多字段问题
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 方法
// 接收者使用对应类型首字母小写
// 指定接收者后，只有接收者这个类型的变量才可以调用此方法
func (p person) dream(str string) {
	fmt.Printf("%s's dream is to learn %s language well!\n", p.name, str)
}

// func (p person) guonian() {
// 	p.age++ // 此处的p是p1的副本，改的是副本
// }

// 指针接收者
// 1. 需要修改结构体变量的值
// 2. 结构体本身较大，拷贝的内存开销比较大
// 3. 保持一致性，如果有方法使用了指针接收者，其他地方的也要用
func (p *person) guonian() { /* person类型的指针 */
	p.age++
}

func main() {
	var s = a{
		10,
		"dsajf",
	}
	// var a = struct {	/* 匿名结构体（无名字结构体） */
	// 	x int
	// 	y string
	// }{10, "ahkd"}	/* “{}”是赋值的花括号 */
	// fmt.Println(a)
	fmt.Println(s)

	// 三种结构体实例化方式：
	var p1 person
	p1.name = "dsjkf"
	p1.age = 100

	p2 := person{"jaf123", 103}

	// 调用构造函数生成person类型变量
	p3 := newPerson("fak", 190)

	fmt.Println(p1, p2, p3)

	p1.dream("c++")
	p2.dream("java")

	fmt.Println(p1.age) /* 值类型 */
	p1.guonian()
	fmt.Println(p1.age)

	// 结构体嵌套
	type addr struct {
		dizhi string
		city  string
	}

	type student struct {
		name    string
		address addr // 嵌套别的结构体
		addr         // 匿名嵌套，只使用类型
	}

	// 序列化
	type point struct {
		// x int
		// y int
		X int    `json:"x"` // 是为了把X转为x，x也可换成任意字符串
		Y string `json:"y"`
	}
	po := point{10371, "1237894"}
	b, err := json.Marshal(po)
	if err != nil {
		fmt.Printf("failed:%v\n", err)
	}
	fmt.Println(string(b)) // 结果为“{}”，因为结构体定义时首字母未大写，访问不到

	// 反序列化
	str := `{"x":19,"y":"asdf17"}`
	var pp point
	// json.Unmarshal([]byte(str), pp)	// 此时修改的是pp的副本

	/* 反序列化时要传递指针！！！ */

	json.Unmarshal([]byte(str), &pp) // 地址引用，修改的是pp
	if err != nil {
		fmt.Printf("failed:%v\n", err)
	}
	fmt.Println(pp)
}
