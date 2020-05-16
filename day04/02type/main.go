package main

import "fmt"

type myint int  //自定义类型
type yint = int //类型别名

func main() {
	var n myint
	// n = 200
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yint
	// m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune //底层rune=int32,不用int32是因为表示变量存的是字符，
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)

}
