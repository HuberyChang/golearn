package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {
	s1 := Student{"几点啦", 80, "放进去"}
	fmt.Printf("%T,%p\n", s1, &s1) //main.Student,0xc000064330
	p1 := &s1
	fmt.Printf("%T\n", p1) //*main.Student
	//fmt.Println((*p1).Name) //语法糖
	fmt.Println(p1.Name)

	//重新赋值
	strPointer := reflect.ValueOf(&s1)
	fmt.Printf("...%T\n", strPointer)                                         //reflect.Value
	fmt.Printf("name:'%v' kind:'%v'\n", strPointer.Type(), strPointer.Kind()) //name:'*main.Student' kind:'ptr'
	typeOfstu := reflect.TypeOf(&s1)                                          //name:'' kind:'ptr'
	//typeOfstu := reflect.TypeOf(p1)
	//反射中对所有指针变量的种类都是 Ptr，但注意，指针变量的类型名称是空，不是 *Student。
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfstu.Name(), typeOfstu.Kind()) //name:'' kind:'ptr'
	if strPointer.Kind() == reflect.Ptr {
		newValue := strPointer.Elem()
		fmt.Printf("%T\n", newValue) //reflect.Value
		fmt.Println(newValue.CanSet())

		f1 := newValue.FieldByName("Name")
		f1.SetString("付南风")
		f3 := newValue.FieldByName("School")
		f3.SetString("暗红色的覅㪐")
		fmt.Println(s1)
	}
}
