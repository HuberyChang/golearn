package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string //字段名之所以大写，是因为结构体中只有可导出的字段是“可设置”的
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello :", msg)
}

func (p Person) personInfo() {
	fmt.Printf("name:%s age:%d sex:%s", p.Name, p.Age, p.Sex)
}

func getMsg(input interface{}) {
	getType := reflect.TypeOf(input)             //先获取类型
	fmt.Println("get Type is :", getType.Name()) //Person
	fmt.Println("get Kind is :", getType.Kind()) // Struct

	getValue := reflect.ValueOf(input)
	fmt.Println("all info:", getValue)

	//获取字段
	/*
		1.先获取Type对象：reflect.Type
			NumField()
			Filed(index)
		2.通过Field()获取每个Field字段
		3.Interface()，得到对应的Value
	*/
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("字段名：%s, 字段类型：%s, 字段值：%v\n", field.Name, field.Type, value)
	}

	//获取方法
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名：%s, 方法类型：%v\n", method.Name, method.Type)
	}
}

func main() {
	p1 := Person{"大", 19, "氨基酸的"}
	getMsg(p1)
}
