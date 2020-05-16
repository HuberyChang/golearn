package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]int, 5, 10) // 创建长度为5，容量为10的切片
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	a1 := [...]int{10, 3, 8, 7, 1}
	sort.Ints(a1[:])	// 切片排序
	fmt.Println(a1)
}
