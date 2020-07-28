package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.23
	fmt.Println(num)

	//修改值
	//对应的要传入的是指针，同时要通过Elem方法获取原始值对应的反射对象
	pointer := reflect.ValueOf(&num)      //必须使用地址。pointer 中的值还是不可取地址,它只是一个指针 &num 的拷贝。所有通过 reflect.ValueOf(x) 返回的 reflect.Value 都是不可取地址的。
	fmt.Printf("pointer类型：%T\n", pointer) //reflect.Value
	fmt.Printf("%p\n", &pointer)          //0xc0000044e0
	newValue := pointer.Elem()            //Elem 方法能够对指针进行“解引用”，来获取任意变量x对应的可取地址的 Value。然后将结果存储到反射 Value 类型对象 newValue 中
	fmt.Println(newValue.Addr())          //0xc00000a0c0
	fmt.Printf("%p\n", &newValue)         //0xc000004520

	fmt.Println("type:", newValue.Type())    //float64
	fmt.Println("canset", newValue.CanSet()) //true

	//重新赋值
	newValue.SetFloat(3.14)
	fmt.Println(num)
}
