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
	s2 = []string{"沙河", "张江", "南山"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s1 == nil)

	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	/*
		数组得到切片
		slice:=arr[start:end]  包左不包右

		从已有的数组上，直接创建切片，该切片的底层数组就是当前的数组
			长度是从start到end的数据量
			容量是从start到数组的结尾
	*/
	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("===========1.已有数组直接创建切片==========")
	s3 := a1[:5] // 包左不包右
	s4 := a1[3:8]
	s5 := a1[5:]
	s6 := a1[:]
	fmt.Println(a1)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)

	fmt.Printf("%p\n", &a1)
	fmt.Printf("%p\n", s3)

	fmt.Println("===========2.长度和容量==========")
	fmt.Printf("len(s3):%d,cap(s3):%d\n", len(s3), cap(s3)) // len(s3):5,cap(s3):10
	fmt.Printf("len(s4):%d,cap(s4):%d\n", len(s4), cap(s4)) // len(s4):5,cap(s4):7
	fmt.Printf("len(s5):%d,cap(s5):%d\n", len(s5), cap(s5)) // len(s5):5,cap(s5):5
	fmt.Printf("len(s6):%d,cap(s6):%d\n", len(s6), cap(s6)) // len(s6):10,cap(s6):10

	fmt.Println("===========3.更改数组的内容==========")
	a1[4] = 100
	fmt.Println(a1)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)

	fmt.Println("===========4.更改切片的内容==========")
	s4[2] = 200
	fmt.Println(a1)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	fmt.Println("===========5.append切片的内容==========")
	s3 = append(s3, 1, 1, 1, 1)
	fmt.Println(a1)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	fmt.Println("===========5.append切片的内容扩容==========")
	fmt.Println(len(s3), cap(s3))
	s3 = append(s3, 2, 2, 2, 2, 2)
	fmt.Println(a1)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(len(s3), cap(s3))
	fmt.Printf("%p\n", s3)
	fmt.Printf("%p\n", &a1)

	/*	// 切片再切片
		s7 := s5[3:]
		fmt.Printf("len(s7):%d cap(s7):%d\n", len(s7), cap(s7))
		// 切片是引用类型，都指向底层的一个数组
		a1[6] = 1300 // 修改底层的数组的值
		s3[2] = 9000
		fmt.Printf("s3:%d s6:%d ", s3, s6)*/
}
