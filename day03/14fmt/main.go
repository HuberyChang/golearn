package main

import "fmt"

func main() {
	// fmt.Println("jkgsd")

	// //整数->字符
	// fmt.Printf("%q\n", 65)

	// //浮点数和复数
	// // fmt.Printf("%b\n", 3.141592653)

	// fmt.Printf("%9.2f", 3.139188)

	// var s string
	// fmt.Scan(&s)
	// fmt.Println("input string：", s)

	var (
		name  string
		age   int
		class string
	)
	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}
