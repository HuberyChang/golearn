package main

import "fmt"

func main() {
	// 切片
	var s1 []int // 定义一个存放int类型元素的切片
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河","张江", "南山"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s1 == nil)

	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))


	// 数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4]	// 包左不包右
	s4 := a1[:4]
	s5 := a1[3:]
	s6 := a1[:]
	fmt.Println(s3, s4, s5, s6)
	// 切片再切片
	s7 := s5[3:]
	fmt.Printf("len(s7):%d cap(s7):%d\n", len(s7),cap(s7))
	// 切片是引用类型，都指向底层的一个数组
	a1[6] = 1300	// 修改底层的数组的值
	s3[2] = 9000
	fmt.Printf("s3:%d s6:%d ", s3, s6)
}
