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

}
