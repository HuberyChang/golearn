package main

import "fmt"

// 整型

func main()  {
	var i1 = 101
	fmt.Printf("%d\n",i1)
	i2 := 077
	fmt.Printf("%d\n",i2)
	i3 := 0x124337843
	fmt.Printf("%d\n",i3)
	// 声明int8类型
	i4 := int8(9)	// 明确类型
	fmt.Printf("%T",i4)
}