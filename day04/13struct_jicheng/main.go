package main

import "fmt"

// 结构体模拟实现其他语言的“继承”

type animal struct {
	name string
}

// 给animal实现一个会移动的方法
func (a animal) move() {
	fmt.Printf("%s会飞\n", a.name)
}

// 狗类

type dog struct {
	feet uint8
	animal
}

// 给狗实现一个汪汪的方法

func (d dog) wang() {
	fmt.Printf("%s会叫\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			name: "haha",
		},
		feet: 4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
