package main

import (
	"fmt"
	"strconv"
)

/*
	type: 用于类型定义和类型别名
	1.类型定义：type 类型名 Type
	2.类型别名：type 类型名 = Type
*/

//1.定义新类型
type myint int //自定义类型
type mysring string
type yint = int //类型别名

//2.定义函数类型
type myfun func(int, int) string

func fun1() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

func main() {
	var n1 myint
	var n2 = 300
	n1 = 200
	fmt.Println(n1, n2)
	fmt.Printf("%T\n", n1)

	var name mysring
	name = "华士大夫"
	var s1 string
	s1 = "氨基酸的"
	fmt.Println(name, s1)
	//n1 = n2 //cannot use n2 (type int) as type myint in assignment
	//name = s1 // cannot use s1 (type string) as type mysring in assignment
	fmt.Printf("%T,%T,%T,%T\n", n1, n2, name, s1) //main.myint,int,main.mysring,string
	fmt.Println("++++++++++++++++++++++++")
	ret := fun1()
	fmt.Println(ret(10, 20))

	var m yint
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	m = n2
	fmt.Println(m)

	var c rune //底层rune=int32,不用int32是因为表示变量存的是字符，
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)

}
