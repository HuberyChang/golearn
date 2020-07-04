package main

import "fmt"

type Person struct {
	name string
}

func (p Person) show() {
	fmt.Println("person------", p.name)
}

type People = Person

func (p People) show2() {
	fmt.Println("people------", p.name)
}

type Student struct {
	//嵌入两个结构体
	People
	Person
}

func main() {
	var s Student
	s.Person.name = "放假啊"
	s.People.name = "就放弃"
	//s.show() //ambiguous selector s.show
	s.Person.show()
	s.People.show()
	s.People.show2()
}
