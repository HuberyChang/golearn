package main

import "fmt"

// copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1	// 赋值
	var a3 = make([]int, 3, 5)
	// a1[0] = 200
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 将索引为1 的删除
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)
	fmt.Println(cap(a1))


	x1 := [...]int{1, 3 ,5, 7, 9, 11, 13, 15}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))
	// 1.切片不保存值
	// 2.切片对应一个底层数组
	// 3.底层数组占用一块连续空间
	s1 = append(s1[:1], s1[2:]...)	// 修改了底层数组！！！
	fmt.Println(s1, len(s1), cap(s1))

	fmt.Println(x1)
}
