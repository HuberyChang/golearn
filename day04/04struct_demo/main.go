package main

import "fmt"

// 结构体是值类型

type people struct {
	name, gender string
}

// go语言中函数传参永远是拷贝
func f(x people) {
	x.gender = "nan" //修改的是副本的gender
}

func f2(x *people) {
	(*x).gender = "nan" //根据内存地址找到那个原变量，修改的就是原来的变量
	// x.gender = "nan" //语法糖，自动根据指针找对应的变量
}

func main() {
	var p people
	p.name = "jkdlfa"
	p.gender = "kjl"
	f(p)
	fmt.Println(p.gender) //kjl
	f2(&p)
	fmt.Println(p.gender)

	var ps = new(people) //对应类型的指针
	// (*ps).gender = "jkfla"
	ps.gender = "jqei" //语法糖
	fmt.Printf("%T\n", ps)
	fmt.Printf("%p\n", ps)  //ps保存的值就是一个内存地址
	fmt.Printf("%p\n", &ps) //&ps求ps的值

	//结构体指针
	// 2.1 k-v初始化
	var p3 = people{
		name:   "9au09",
		gender: "nand",
	}

	// var p3 = &people{	//不用new，直接生成一个people类型的指针
	// 	name:   "9au09",
	// 	gender: "nand",
	// }
	fmt.Printf("%v\n", p3)

	// 2.2、使用值列表的形式初始化，值的顺序要和结构体定义时字段顺序一致
	p4 := people{
		"华科技的说法",
		"拉屎",
	}

	// p4 := &people{	//不用new，直接生成一个people类型的指针
	// 	"华科技的说法",
	// 	"拉屎",
	// }

	fmt.Printf("%v\n", p4)
}
