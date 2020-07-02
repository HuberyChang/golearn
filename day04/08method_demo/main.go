package main

import "fmt"

// 自定义类型加方法
// 只能给自己包里定义的类型添加方法，不能给别的包里面的类型添加方法
type myInt int

type Worker struct {
	name string
	age  int
	sex  string
}

type Cat struct {
	color string
	age   int
}

func (w Worker) work() { // w = w1
	fmt.Println(w.name, "找工作。。。")
}

func (w *Worker) rest() { // p=w2,p=w1的地址
	fmt.Println(w.name, "在写啊客户端借款方。。。")
}

func (w *Worker) printInfo() {
	fmt.Printf("工人姓名：%s 工人年龄：%d 工人性别：%s\n", w.name, w.age, w.sex)
}
func (c *Cat) printInfo() {
	fmt.Printf("猫咪颜色：%s 猫咪年龄：%d\n", c.color, c.age)
}

func (i myInt) hello() {
	fmt.Println("int")
}

func main() {
	m := myInt(100)
	m.hello()

	w1 := Worker{name: "爱的世界", age: 18, sex: "难"}
	fmt.Printf("%T\n", w1) //main.Worker
	w1.work()

	w2 := &Worker{name: "qui家居", age: 19, sex: "打飞机"}
	fmt.Printf("%T\n", w2) //*main.Worker
	w2.work()

	w2.rest()
	w1.rest()

	w2.printInfo()
	c1 := Cat{color: "black", age: 9}
	c1.printInfo()
}
