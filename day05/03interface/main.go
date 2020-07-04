package main

import "fmt"

// 引出接口
// 接口就是一种类型

// 定义一个能叫的变量
type speaker interface {
	speak() //只要实现speakf方法的都是speaker类型
}

type cat struct {
}

type dog struct {
}

type person struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("哈哈哈")
}

func da(x speaker) { //接受一个参数
	x.speak()
	// fmt.Printf("%T\n", x)
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person
	// fmt.Printf("%T\n", c1)
	/*
		当需要接口类型的对象时，可以使用任意实现类对象替代：c1/d1/p1
	*/
	da(c1)
	da(d1)
	da(p1)

	var sp speaker //定义一个speaker类型的变量
	/*
		接口对象不能访问实现类中的属性
	*/
	sp = c1 // c1可以访问cat结构体的内部字段，但是sp不能访问
	sp.speak()
	fmt.Println(sp)
}
