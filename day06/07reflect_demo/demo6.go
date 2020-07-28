package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func f1() {
	fmt.Println("fun1,无参数。。。")
}

func f2(i int, s string) {
	fmt.Println("fun2,有参数。。。")
}

func f3(i int, s string) string {
	fmt.Println("fun3,有参数和返回值。。。", i, s)
	return s + strconv.Itoa(i)
}

func main() {
	/*
		思路：函数也可以看做是接口类型变量
		1.函数--->反射对象，value
		2.kind--->func
		3.call()
	*/
	f1 := f1
	value := reflect.ValueOf(f1)
	fmt.Printf("kind:%s type:%s\n", value.Kind(), value.Type()) //kind:func type:func()

	value2 := reflect.ValueOf(f2)
	fmt.Printf("kind:%s type:%s\n", value2.Kind(), value2.Type()) //kind:func type:func(int, string)

	value3 := reflect.ValueOf(f3)
	fmt.Printf("kind:%s type:%s\n", value3.Kind(), value3.Type()) //kind:func type:func(int, string) string

	//通过反射调用
	value.Call(nil)
	value2.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf("SDK浪费")})
	ret := value3.Call([]reflect.Value{reflect.ValueOf(10), reflect.ValueOf("嗷嗷的就")})
	fmt.Printf("%T\n", ret)                                        //[]reflect.Value
	fmt.Println(len(ret))                                          //1
	fmt.Printf("kind:%s, type:%s\n", ret[0].Kind(), ret[0].Type()) //kind:string, type:string

	s := ret[0].Interface().(string)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}
