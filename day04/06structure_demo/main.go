package main

import "fmt"

// 构造函数：返回一个结构体变量的函数
type person struct {
	name string
	age  int
}

// 构造函数，约定成俗以new开头
// 返回的是结构体还是结构体指针？
// 结构体是值类型，赋值的时候是拷贝，结构体比较小的时候，返回结构体
// 结构体比较大的时候，尽量返回结构体指针，减少程序的内存开销
func newPerson(name string, age int) person {
	return person{ //返回结构体
		name: name,
		age:  age,
	}
}

// func newPerson(name string, age int) person {
// 	return &person{	//返回结构体指针
// 		name: name,
// 		age:  age,
// 	}
// }

func main() {
	p1 := newPerson("花甲粉", 9)
	p2 := newPerson("全陪团", 10)
	fmt.Println(p1, p2)
}
