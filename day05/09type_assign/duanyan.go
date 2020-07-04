package main

import (
	"fmt"
	"math"
)

/*
	接口断言：
		方式一：
			1. instance := 接口对象.(实际类型)  // 不安全，会panic()
			2. instance,ok := 接口对象.(实际类型)  // 安全
		方式二：
			switch instance := 接口对象.(type){
			case 实际类型1:
				...
			case 实际类型2:
				...
			...
			}
*/

//定义一个接口
type Shape interface {
	peri() float64 // 周长
	area() float64 // 面积
}

type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

type Circle struct {
	r float64
}

func (c Circle) peri() float64 {
	return 2 * c.r * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.r, 2) * math.Pi
}

func testSharp(s Shape) {
	fmt.Printf("周长：%.2f,面积：%.2f\n", s.peri(), s.area())
}

func getType(s Shape) {
	//断言
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三边是：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("圆形，圆半径：", ins.r)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Println("是三角形，三边是：", ins.a, ins.b, ins.c)
		fmt.Printf("ins:%T,%p,%p\n", ins, &ins, ins) // ins:*main.Triangle,0xc000006038,0xc00000c440
		fmt.Printf("s:%T,%p,%p\n", s, &s, s)         // s:*main.Triangle,0xc000030220,0xc00000c440
	} else {
		fmt.Println("不晓得")
	}
}

func getType2(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形:", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形：", ins.r)
	case *Triangle:
		fmt.Println("三角形指针：", ins.a, ins.b, ins.c)
	}
}

func main() {
	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1)

	var c1 Circle = Circle{2}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1)

	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	var s2 Shape
	fmt.Printf("s2:%T,%p,%p\n", s2, &s2, s2)
	s2 = c1
	fmt.Println(s2.peri())
	fmt.Println(s2.area())

	testSharp(t1)
	testSharp(c1)
	testSharp(s1)

	getType(t1)
	getType(c1)
	getType(s1)

	var t2 *Triangle = &Triangle{3, 4, 2}
	fmt.Printf("t2:%T,%p,%p\n", t2, &t2, t2) // t2:*main.Triangle,0xc000006030,0xc00000c440
	getType(t2)

	getType2(t2)
	getType2(t1)
}
