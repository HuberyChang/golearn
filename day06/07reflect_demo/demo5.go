package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string
	Age  int
	Sex  string
}

func (p People) Say(msg string) {
	fmt.Println(msg)
}

func (p People) PrintInfo() {
	fmt.Printf("name:%s, age:%d, sex:%s\n", p.Name, p.Age, p.Sex)
}

func (p People) Test(i, j int, s string) {
	fmt.Println(i, j, s)
}

func main() {
	p1 := People{"罚款", 678, "大家是否"}
	value := reflect.ValueOf(p1)
	fmt.Printf("kind:%s,type:%s\n", value.Kind(), value.Type()) //kind:struct,type:main.People

	methodv1 := value.MethodByName("PrintInfo")
	fmt.Printf("kind:%s, type:%s\n", methodv1.Kind(), methodv1.Type()) //kind:func, type:func()

	//没有参数，进行调用
	//一：用nil
	methodv1.Call(nil)
	//二：用空切片
	args := make([]reflect.Value, 0)
	methodv1.Call(args)

	//有参数，进行调用
	methodv2 := value.MethodByName("Say")
	fmt.Printf("kind:%s, type:%s\n", methodv2.Kind(), methodv2.Type()) //kind:func, type:func(string)
	args1 := []reflect.Value{reflect.ValueOf("反射机制")}
	methodv2.Call(args1)

	methodv3 := value.MethodByName("Test")
	fmt.Printf("kind:%s, type:%s\n", methodv3.Kind(), methodv3.Type()) //kind:func, type:func(int, int, string)
	args3 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200), reflect.ValueOf("发虹口区开发")}
	methodv3.Call(args3)

}
