package main

import "fmt"

/*
	append为切片追加元素
	1、向切片中添加数据时，如果没有超过容量，则直接添加，如果超过容量，自动扩容（修改底层数组地址，同时容量成倍扩容）
	2、切片一旦扩容，就是重新指向一个新的底层数组
*/
func main() {
	s1 := []string{"北京", "上海", "成都"}
	fmt.Printf("%p\n", s1)                                                         // 0xc0000b4330
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d   ==========\n", s1, len(s1), cap(s1)) // len(s1):3 cap(s1):3

	// s1[3] = "广州"	// 错误写法，索引越界
	// 调用append函数必须使用原来切片的变量接收返回值
	s1 = append(s1, "广州") //append 函数追加元素，原来的底层数组放不下的时候，go底层就会把底层数组换一个
	fmt.Println(s1)
	fmt.Printf("%p\n", s1)                                                         // 0xc0000d4000
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d   ==========\n", s1, len(s1), cap(s1)) // len(s1):4 cap(s1):6

	s1 = append(s1, "深圳")
	fmt.Println(s1)
	fmt.Printf("%p\n", s1)                                                         // 0xc0000d4000
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d   ==========\n", s1, len(s1), cap(s1)) // len(s1):5 cap(s1):6

	ss := []string{"武汉", "苏州"}
	s1 = append(s1, ss...) // ...表示拆开
	fmt.Println(s1)
	fmt.Printf("%p\n", s1)                                                         // 0xc0000d8000
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d   ==========\n", s1, len(s1), cap(s1)) // len(s1):7 cap(s1):12
}
