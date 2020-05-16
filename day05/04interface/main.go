package main

import "fmt"

type fll struct {
	brand string
}

type bsj struct {
	brand string
}

// 定义一个car接口
// 不管什么结构体，只要run方法都能是car类型
type car interface {
	run()
}

func (f fll) run() {
	fmt.Printf("%s速度70迈\n", f.brand)
}

func (b bsj) run() {
	fmt.Printf("%s速度700迈\n", b.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	var f = fll{
		brand: "法拉利",
	}
	var b = bsj{
		brand: "保时捷",
	}
	drive(f)
	drive(b)
}
