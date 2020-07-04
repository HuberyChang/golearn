package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet string
}

type A interface {
	testA()
}

type B interface {
	testB()
}

type C interface {
	A
	B
	testC()
}

type Cat struct {
}

func (c Cat) testA() {
	fmt.Println("testA")
}
func (c Cat) testB() {
	fmt.Println("testB")
}
func (c Cat) testC() {
	fmt.Println("testC")
}

// cat实现了mover接口
func (c *cat) move() {
	fmt.Println("走猫步")
}

// cat实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("吃%s时", food)
}

// func (c *cat) move() {
// 	fmt.Println("猫走步")
// }

// func (c *cat) eat(food string) {
// 	fmt.Println("吃%s是", food)
// }

func main() {
	var cat Cat = Cat{}
	cat.testA()
	cat.testB()
	cat.testC()
	fmt.Println("=++++++++++++=")
	var a1 A = cat
	a1.testA()
	fmt.Println("=++++++++++++=")
	var b1 B = cat
	b1.testB()
	fmt.Println("=++++++++++++=")
	var c1 C = cat
	c1.testA()
	c1.testB()
	c1.testC()
	fmt.Println("=++++++++++++=")
	//var c2 C=a1 // 不行
	var a2 A = c1
	a2.testA() // 只能实现testA方法
}
