package main

import "fmt"

// make 函数创造切片

func main() {
	s1 := make([]int, 5, 10)	// 底层数组容量10，切+ 片长度为5
	s2 := make([]int, 5)		// 底层数组容量为5，切片长度为5
	fmt.Printf("s1:%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))

	// 切片赋值
	s3 := []int{1, 3, 5}
	s4 := s3	// s3, s4都指向同一个底层数组
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	// 切片的遍历
	// 1.索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(i, s3[i])
	}

	// 2.for range 循环
	for i, v := range s3 {
		fmt.Println(i, v)
	}
}
