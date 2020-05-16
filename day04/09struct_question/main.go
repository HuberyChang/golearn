package main

import "fmt"

// 结构体遇到的问题
// 1.myInt(100)是啥？类比int32(100)

type myInt int

func (m myInt) hello() {
	fmt.Println("int")
}

func main() {
	// 声明一个int32类型的变量，它的值是10
	// 方法1：
	// 	var x int32
	// 	x=10
	// 方法2：
	// 	var x int32 = 10
	// 方法3：
	// 	var x = int32(10)
	// 方法4：
	// 	x:=int32(10)
	// fmt.Println(x)

	// 声明一个myInt类型的变量，它的值是10
	// 方法1：
	// 	var x myInt
	// 	x=10
	// 方法2：
	// 	var x myInt = 10
	// 方法3：
	// 	var x = myInt(10)
	// 方法4：
	// 	x:=myInt(10)	//强制类型转换
	// fmt.Println(x)

	// 结构体初始化
	type person struct {
		name string
		age  int
	}

	// 方法1：声明变量再初始化
	var p person
	p.name = "kjfal"
	p.age = 19
	fmt.Println(p)

	// 方法2：
	// 2.1：k-v初始化，声明同时初始化
	var p1 = person{
		name: "hfa",
		age:  19,
	}
	fmt.Println(p1)

	// 2.2：值列表初始化
	var p2 = person{
		"hfas",
		18,
	}
	fmt.Println(p2)

}

type person struct {
	name string
	age  int
}

// 为什么要用构造函数
func newPerson(name string, age int) person { //返回一个person类型的变量，根据person结构体的定义，需要传递两个参数，参数名称任意定义，但是最好跟结构体定义一样
	// 别人调用我，我可以给他一个person类型的变量

	// 返回值类型
	// 1.
	// tmp:=person{
	// 	name:name,
	// 	age:age,
	// }
	// return tmp
	// 2.
	return person{
		name: name,
		age:  age,
	}

}

/* func newPerson(name string, age int) *person { //返回一个person类型的变量，根据person结构体的定义，需要传递两个参数，参数名称任意定义，但是最好跟结构体定义一样
	// 别人调用我，我可以给他一个person类型的变量

	// 返回指针类型
	return &person{
		name: name,
		age:  age,
	}

} */
