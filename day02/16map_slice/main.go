package main

import "fmt"

// map和slice组合
func main() {
	// 元素类型为map的切片
	// 对map和slice同时初始化
	var s1 = make([]map[int]string, 1, 10)	// 初始化slice
	// 没对内部的map做初始化
	// s1[0][100] = "A"
	s1[0] = make(map[int]string, 1)	// 初始化map
	s1[0][100] = "hang"
	fmt.Println(s1)



	// 值为slice类型的map
	var m1 = make(map[string][]int, 10)
	m1["沙河"] = []int{10, 20 ,30}
	fmt.Println(m1)
}
