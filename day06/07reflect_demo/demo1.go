package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.23

	//接口类型变量--->反射类型对象
	value := reflect.ValueOf(num)
	fmt.Println(value)
	fmt.Printf("%T\n", value) //reflect.Value
	typen := reflect.TypeOf(num)
	fmt.Printf("%T\n", typen) //*reflect.rtype
	fmt.Println(typen)        //float64

	//反射类型对象--->接口类型变量
	conValue := value.Interface().(float64) //获取interface{}类型的值, 通过断言，恢复底层的具体值
	fmt.Printf("%T\n", conValue)
	fmt.Println(conValue)

	/*
		反射类型对象--->接口类型变量,强制装换
		float64，*float64不一样，弄错会panic
	*/
	//pointer:=value.Interface().(float64)
	pointer := reflect.ValueOf(&num)
	conPointer := pointer.Interface().(*float64)
	fmt.Println(conPointer) //0xc00000a0c0

}
