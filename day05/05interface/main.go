package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(string)
	// eat()
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c cat) move() {
	fmt.Println("miao:喊")
}

func (c chicken) move() {
	fmt.Println("鸡冻")
}

func (c cat) eat(x string) {
	fmt.Printf("喵%s\n", x)
}

func (c chicken) eat(y string) {
	fmt.Printf("鸡%s\n", y)
}

func main() {
	var a1 animal //定义一个接口类型

	bc := cat{
		name: "hsadjk",
		feet: 4,
	}
	a1 = bc
	a1.eat("鲸鲨")
	fmt.Println(a1)
	ji := chicken{
		feet: 2,
	}
	a1 = ji
	fmt.Println(a1)
	a1.eat("沙发哈")
	fmt.Println(a1)
}
