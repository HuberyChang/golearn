package main

import "fmt"

func main() {
	// 1.取地址
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)

	
	// 2.根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)



	// var a *int	// nil pointer
	// fmt.Println(a)
	// *a = 100
	// fmt.Println(*a)

	// new函数申请内存地址
	fmt.Println("========================")
	var a = new(int)	
	fmt.Println(a)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)
	fmt.Println("========================")

	// var b map[string]int
	// b["沙河扛把子"] = 100
	// fmt.Println(b)
}
