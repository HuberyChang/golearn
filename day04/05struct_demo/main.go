package main

import "fmt"

// 结构体占用一块连续的内存空间

type x struct {
	a int8
	b int8
	c int8
}

type Book struct {
	bkname string
	price  float64
}

type Student struct {
	name string
	age  int
	book Book
}

type Student1 struct {
	name string
	age  int
	book *Book
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))

	b := Book{}
	b.bkname = "大师傅"
	b.price = 19.78
	s := Student{}
	s.name = "取得了父母"
	s.age = 19
	s.book = b
	fmt.Println(b) //{大师傅 19.78}
	fmt.Println(s) //{取得了父母 19 {大师傅 19.78}}

	b1 := Book{"橘子洲头", 189.90}
	s1 := Student1{"就放弃搭建", 89, &b1}
	fmt.Println(b1) //{橘子洲头 189.9}
	fmt.Println(s1) //{就放弃搭建 89 0xc000004500}
	fmt.Println(s1.book)

	s1.book.bkname = "西游记"
	fmt.Println(b1) //{西游记 189.9}
	fmt.Println(s1) //{就放弃搭建 89 0xc000004500}

	s2 := Student{
		name: "放弃",
		age:  19,
		book: Book{
			"放进去",
			19.678,
		},
	}
	fmt.Println(s2)

	s3 := Student{
		"没看到钱",
		197,
		Book{
			"开发区",
			190.08,
		},
	}
	fmt.Println(s3)

}
